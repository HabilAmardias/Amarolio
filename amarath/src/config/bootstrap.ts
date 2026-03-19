import Elysia from "elysia";
import { AppRouter } from "../routers/AppRouter";
import { NewMessageService } from "../services/MessageService";
import { RedisClient, SQL } from "bun";
import { NewOllamaUtil } from "../utils/Ollama";
import { NewMessageController } from "../controller/MessageController";
import { NewMessagePublisher } from "../repository/redis/MessagePublisher";
import { NewChatroomService } from "../services/ChatroomService";
import { NewChatroomController } from "../controller/ChatroomController";
import { Logger } from "@logtape/logtape";

export function Bootstrap(
  app: Elysia,
  redisClient: RedisClient,
  pgClient: SQL,
  logger: Logger,
) {
  const ollamaUtil = NewOllamaUtil();

  const messagePublisherRepo = NewMessagePublisher(redisClient);

  const messageService = NewMessageService(
    messagePublisherRepo,
    ollamaUtil,
    pgClient,
  );

  const chatroomService = NewChatroomService(
    pgClient,
    ollamaUtil,
    messagePublisherRepo,
  );
  const messageController = NewMessageController(messageService);
  const chatroomController = NewChatroomController(chatroomService);

  const appRouter = new AppRouter(
    messageController,
    chatroomController,
    logger,
    app,
  );

  appRouter.Setup();
}
