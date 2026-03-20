export const ProductionMode = "PRODUCTION";

export const ContextPrompt = (ctx: string) => {
  return `You are given external web search context.
Use ONLY this context to answer.
If the answer is not contained, say you don't know.

<context>
${ctx}
</context>
`;
};
