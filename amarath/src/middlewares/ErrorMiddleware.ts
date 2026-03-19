import Elysia from "elysia";
import { CustomError, InvalidActionError, CommonError } from "../customerror";
import { Logger } from "@logtape/logtape";

export function NewErrorMiddleware(logger: Logger) {
  return new Elysia()
    .error({ CustomError })
    .onError({ as: "global" }, ({ code, error, request, path, set }) => {
      let err: CustomError;
      switch (code) {
        case "CustomError":
          err = error;
          break;
        case "VALIDATION":
          err = new CustomError(
            error.message,
            InvalidActionError,
            error.detail(error.message).toString(),
          );
          break;
        default:
          err = new CustomError(
            "something went wrong",
            CommonError,
            (error as Readonly<Error>).message,
          );
          break;
      }
      logger.error(
        new Error(
          `${request.method} ${path}: ${set.status} - ${err.sysMessage}`,
        ),
      );
      return err.toResponse();
    });
}
