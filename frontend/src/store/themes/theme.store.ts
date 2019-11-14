import { Module } from "vuex";
import { RootState } from "..";
import { actions } from "./theme.actions";
import { mutations } from "./theme.mutations";

export const namespaced: boolean = true;
export const namespace: string = "themes";

export interface IThemeState {
  themes: string[];
  selected?: string;
  uploadedFile?: string;
}

export const state: IThemeState = {
  themes: [],
  selected: undefined,
  uploadedFile: undefined
};

export const themes: Module<IThemeState, RootState> = {
  state,
  actions,
  mutations,
  namespaced
};
