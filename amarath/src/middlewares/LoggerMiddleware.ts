import { Logger } from "@logtape/logtape";
import Elysia from "elysia";

export function NewLoggerMiddleware(logger: Logger) {
  return new Elysia()
    .resolve({ as: "global" }, () => {
      const start = performance.now();
      return { start };
    })
    .onAfterHandle({ as: "global" }, ({ request, path, start, set }) => {
      const duration = performance.now() - start;
      logger.info(
        `${request.method} ${path}: ${set.status} - ${duration.toFixed(2)}ms`,
      );
    });
}
