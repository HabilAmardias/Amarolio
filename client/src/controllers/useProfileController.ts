import { useAtom } from "jotai";
import { profileAtom } from "../store/atoms";

export function useProfileController() {
  const [profile] = useAtom(profileAtom);
  return { profile };
}
