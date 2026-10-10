import { atom } from "jotai";
import type { ColorMode } from "../theme/theme";

export const themeModeAtom = atom<ColorMode>("system");
