import { atom } from "jotai";
import type { Profile, Experience, Project } from "../models/types";
import { PROFILE, EXPERIENCES, PROJECTS } from "../models/data";

export const profileAtom = atom<Profile>(PROFILE);
export const experiencesAtom = atom<Experience[]>(EXPERIENCES);
export const projectsAtom = atom<Project[]>(PROJECTS);
export const activeFilterAtom = atom<string>("all");
