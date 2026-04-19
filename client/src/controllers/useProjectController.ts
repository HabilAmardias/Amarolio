import { useAtom } from "jotai";
import { projectsAtom, activeFilterAtom } from "../store/atoms";

export function useProjectController() {
  const [projects] = useAtom(projectsAtom);
  const [activeFilter, setActiveFilter] = useAtom(activeFilterAtom);

  const allTags = [
    "all",
    ...Array.from(new Set(projects.flatMap((p) => p.tags))),
  ];

  const filtered =
    activeFilter === "all"
      ? projects
      : projects.filter((p) => p.tags.includes(activeFilter));

  return { projects: filtered, allTags, activeFilter, setActiveFilter };
}
