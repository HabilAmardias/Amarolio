export class CustomError extends Error {
  status: number;
  sysMessage: string;

  constructor(message: string, errCode: number, sysMessage: string) {
    super(message);
    this.status = Math.round(errCode / 100);
    this.sysMessage = sysMessage;
  }
  toResponse() {
    return Response.json(
      {
        success: false,
        data: {
          message: this.message,
        },
      },
      { status: this.status },
    );
  }
}

export const CommonError = 50001;
export const DatabaseExecError = 50002;
export const ItemNotFoundError = 40401;
export const InvalidActionError = 40001;
export const UnauthorizedError = 40101;
export const TimeoutError = 40801;
