import { file } from "bun";
export const ProductionMode = "PRODUCTION";

export const SystemPrompt = await file(
  `${import.meta.dir}/../prompts/system.md`,
).text();

export const RoutingPrompt = await file(
  `${import.meta.dir}/../prompts/routing.md`,
).text();

export const TitlePrompt = await file(
  `${import.meta.dir}/../prompts/title.md`,
).text();

export const ContextPrompt = (ctx: string) => {
  return `You are given external web search context.
Use ONLY this context to answer.
If the answer is not contained, say you don't know.

<context>
${ctx}
</context>
`;
};
