import { MutationTree } from "vuex";
import { IThemeState } from "./theme.store";

export const mutationSetThemes: string = "setThemes";
export const mutationSetSelectedTheme: string = "setSelectedTheme";
export const mutationSetUploadedFile: string = "setUploadedFile";

export const mutations: MutationTree<IThemeState> = {
  [mutationSetThemes](state: IThemeState, themes: string[]) {
    state.themes = themes;
  },
  [mutationSetSelectedTheme](state: IThemeState, themeId: string) {
    state.selected = themeId;
  },
  [mutationSetUploadedFile](state: IThemeState, fileId: string) {
    state.uploadedFile = fileId;
  }
};
