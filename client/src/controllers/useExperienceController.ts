import { useAtom } from "jotai";
import { experiencesAtom } from "../store/atoms";

export function useExperienceController() {
  const [experiences] = useAtom(experiencesAtom);
  return { experiences };
}
