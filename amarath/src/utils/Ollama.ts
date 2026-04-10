import { Message, Ollama } from "ollama";
import { ContextPrompt } from "../constants";
import { CustomError, TimeoutError } from "../customerror";
import { Glob, file } from "bun";

function ReadPrompts() {
  const glob = new Glob("./prompts/*.md");
  const files = Array.from(glob.scanSync()).sort();
  const promises = files.map((val) => {
    return file(val).text();
  });

  return Promise.all(promises);
}

const [routingPrompt, systemPrompt, titlePrompt] = await ReadPrompts();

export function NewOllamaUtil() {
  return new OllamaUtil();
}

class OllamaUtil {
  ollamaClient: Ollama;
  constructor() {
    this.ollamaClient = new Ollama({
      host: process.env.OLLAMA_HOST,
      headers: {
        Authorization: `Bearer ${process.env.OLLAMA_API_KEY}`,
      },
    });
  }
  CreateTitle = async (userMessage: string) => {
    const MAX_PREDICT = 512;
    const titleMessages = [
      {
        role: "system",
        content: titlePrompt,
      },
      {
        role: "user",
        content: userMessage,
      },
    ];
    let response = await this.ollamaClient.chat({
      model: process.env.TITLE_MODEL_VARIANT || "qwen3:0.6b",
      messages: titleMessages,
      think: false,
      stream: false,
      options: {
        num_predict: MAX_PREDICT,
      },
    });
    return response.message.content;
  };
  WebSearch = async (userMessage: string) => {
    let context = "";
    const routingMessages = [
      {
        role: "system",
        content: routingPrompt,
      },
      {
        role: "user",
        content: userMessage,
      },
    ];
    let routing = await this.ollamaClient.chat({
      model: process.env.MODEL_VARIANT || "qwen3:0.6b",
      messages: routingMessages,
      think: false,
      stream: false,
      options: {
        temperature: 0,
      },
    });

    if (routing.message.content.trim().toUpperCase().startsWith("YES")) {
      const webSearchResponse = await this.ollamaClient.webSearch({
        query: userMessage,
        maxResults: 5,
      });
      context = webSearchResponse.results
        .map((e) => e.content.slice(0, 800))
        .join("\n==========================\n");
    }
    return context;
  };

  SendMessage = async (
    userMessage: string,
    context: string,
    historyMessage?: Message[],
    previousMessage?: string,
    controller?: { signal: AbortSignal },
  ) => {
    const TEMP = 0.1; // low temp for more deterministic answer
    const messages: Message[] = [
      {
        role: "system",
        content: systemPrompt,
      },
    ];
    if (context.length > 0) {
      messages.push({
        role: "system",
        content: ContextPrompt(context),
      });
    }
    if (historyMessage && historyMessage.length > 0) {
      messages.push(...historyMessage);
    }
    messages.push({
      role: "user",
      content: userMessage,
    });
    if (previousMessage) {
      messages.push(
        { role: "assistant", content: previousMessage },
        {
          role: "user",
          content:
            "Continue the explanation from the last incomplete step. Do not repeat earlier steps.",
        },
      );
    }
    let response = await this.ollamaClient.chat({
      model: process.env.MODEL_VARIANT || "qwen3:0.6b",
      messages: messages,
      think: true,
      stream: true,
      tools: [
        // fake function to prevent overthinking
        {
          type: "function",
          function: {
            name: "x",
          },
        },
      ],
      options: {
        temperature: TEMP,
      },
    });
    if (controller?.signal.aborted) {
      throw new CustomError(
        "request timeout",
        TimeoutError,
        "request timeout, operation aborted",
      );
    }
    return response;
  };
}
