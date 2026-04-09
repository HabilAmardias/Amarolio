import { SendMessageParam } from "../entity/MessageEntity";

interface MessageServiceItf {
  SendMessage: (param: SendMessageParam) => Promise<string>;
}

interface MessageReceiverItf {
  ReceiveMessage: (
    group: string,
    stream: string,
    block?: string,
    count?: string,
  ) => Promise<[string, [string, string[]][]][] | null>;
  AcknowledgeMessage: (
    id: string,
    stream: string,
    group: string,
  ) => Promise<void>;
  EnsureGroup: (stream: string, group: string) => Promise<void>;
}

function parseFields(rawFields: string[]): Record<string, string> {
  const fields: Record<string, string> = {};
  for (let i = 0; i < rawFields.length; i += 2) {
    fields[rawFields[i]] = rawFields[i + 1];
  }
  return fields;
}

export async function MessageEventReceiver(
  ms: MessageServiceItf,
  mr: MessageReceiverItf,
) {
  await mr.EnsureGroup("streams:chatrooms", "receiver_worker");
  while (true) {
    const res = await mr.ReceiveMessage("receiver_worker", "streams:chatrooms");
    if (!res) {
      continue;
    }
    for (const [_stream, messages] of res) {
      for (const [id, rawfields] of messages) {
        const fields = parseFields(rawfields);
        const { action, user, chatroom, message } = fields;
        if (action !== "send_message") {
          continue;
        }
        const param: SendMessageParam = {
          userID: user,
          id: chatroom,
          userMessage: message,
        };
        await ms.SendMessage(param);
        await mr.AcknowledgeMessage(id, "streams:chatrooms", "receiver_worker");
      }
    }
  }
}
