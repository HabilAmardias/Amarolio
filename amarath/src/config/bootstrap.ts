import Elysia from "elysia";
import { AppRouter } from "../routers/AppRouter";
import { NewMessageService } from "../services/MessageService";
import { SQL } from "bun";
import { NewOllamaUtil } from "../utils/Ollama";
import { NewMessageController } from "../controller/MessageController";
import { NewMessageBridge } from "../repository/redis/MessageBridge";
import { NewChatroomService } from "../services/ChatroomService";
import { NewChatroomController } from "../controller/ChatroomController";
import { Logger } from "@logtape/logtape";
import { NewChatroomCache } from "../repository/redis/ChatroomCache";
import { NewMessageCache } from "../repository/redis/MessageCache";
import { MessageEventReceiver } from "./worker";
import Redis from "ioredis";

export function Bootstrap(
  app: Elysia,
  redisClient: Redis,
  redisPubSub: Redis,
  pgClient: SQL,
  logger: Logger,
) {
  const ollamaUtil = NewOllamaUtil();

  const messagePublisherRepo = NewMessageBridge(redisPubSub);
  const chatroomCacheRepo = NewChatroomCache(redisClient);
  const messageCacheRepo = NewMessageCache(redisClient);

  const messageService = NewMessageService(
    messagePublisherRepo,
    ollamaUtil,
    chatroomCacheRepo,
    messageCacheRepo,
    pgClient,
  );

  const chatroomService = NewChatroomService(
    pgClient,
    ollamaUtil,
    messagePublisherRepo,
    chatroomCacheRepo,
    messageCacheRepo,
  );
  const messageController = NewMessageController(messageService);
  const chatroomController = NewChatroomController(chatroomService);

  const appRouter = new AppRouter(
    messageController,
    chatroomController,
    logger,
    app,
  );

  MessageEventReceiver(messageService, messagePublisherRepo);

  appRouter.Setup();
}
