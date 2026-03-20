import Elysia from "elysia";
import { Bootstrap } from "./bootstrap";
import { NewRedisClient } from "../db/redis";
import { Migrate, NewPostgresClient } from "../db/postgresql";
import { ProductionMode } from "../constants";
import { configure, getConsoleSink, getLogger } from "@logtape/logtape";

export async function Run() {
  try {
    await configure({
      sinks: { console: getConsoleSink() },
      loggers: [
        { category: "Amarath", lowestLevel: "debug", sinks: ["console"] },
      ],
    });
    const logger = getLogger("Amarath");
    const app = new Elysia({
      cookie: {
        secure: process.env.ENVIRONMENT_MODE === ProductionMode,
        httpOnly: true,
      },
    });
    const redisClient = await NewRedisClient();
    const pgClient = await NewPostgresClient();
    await Migrate(pgClient);

    Bootstrap(app, redisClient, pgClient, logger);

    app.listen({
      port: Number(process.env.SERVER_PORT),
    });
  } catch (err) {
    const castedErr = err as Error;
    console.error(castedErr.name, castedErr.message);
  }
}
