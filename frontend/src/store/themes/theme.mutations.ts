import { MutationTree } from "vuex";
import { IThemeState } from "./theme.store";

export const mutationSetThemes: string = "setThemes";

export const mutations: MutationTree<IThemeState> = {
  [mutationSetThemes](state: IThemeState, themes: string[]) {
    state.themes = themes;
  }
};
