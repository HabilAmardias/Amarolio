import { atom, type PrimitiveAtom } from "jotai";
import type { User } from "./type";

class UserModel {
    private _userAtom: PrimitiveAtom<User>;
    constructor() {
        const initialUser: User | null = null;
        this._userAtom = atom<User | null>(initialUser);
    }

    get userAtom() {
        return this._userAtom;
    }
}

export const userModel = new UserModel()