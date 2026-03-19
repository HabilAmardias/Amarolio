import Elysia, { AnyElysia } from "elysia";
import { Logger } from "@logtape/logtape";
import { NewLoggerMiddleware } from "../middlewares/LoggerMiddleware";
import { NewErrorMiddleware } from "../middlewares/ErrorMiddleware";
import cors from "@elysiajs/cors";

export class AppRouter {
  App: Elysia;
  messageController: AnyElysia;
  chatroomController: AnyElysia;
  logger: Logger;

  constructor(
    messageController: AnyElysia,
    chatroomController: AnyElysia,
    logger: Logger,
    app: Elysia,
  ) {
    this.App = app;
    this.logger = logger;
    this.messageController = messageController;
    this.chatroomController = chatroomController;
  }

  Setup() {
    this.App.use(
      cors({
        origin: process.env.CORS_ORIGIN ? process.env.CORS_ORIGIN : true,
        credentials: true,
        methods: ["GET", "POST", "DELETE"],
      }),
    )
      .use(NewLoggerMiddleware(this.logger))
      .use(NewErrorMiddleware(this.logger))
      .use(this.messageController)
      .use(this.chatroomController);
  }
}
