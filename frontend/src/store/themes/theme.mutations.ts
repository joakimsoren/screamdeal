import { MutationTree } from "vuex";
import { IThemeState } from "./theme.store";

export const mutationSetThemes: string = "setThemes";
export const mutationSetSelectedTheme: string = "setSelectedTheme";
export const mutationSetUploadedFile: string = "setUploadedFile";
export const mutationSetThemedPdf: string = "setThemed";
export const mutationSetLoadingPdf: string = "setLoadingPdf";
export const mutationSetLoadedPdf: string = "setLoadedPdf";
export const mutationSetLoadingApplyTheme: string = "setLoadingApplyTheme";
export const mutationSetLoadedApplyTheme: string = "setLoadedApplyTheme";

export const mutations: MutationTree<IThemeState> = {
  [mutationSetThemes](state: IThemeState, themes: string[]) {
    state.themes = themes;
  },
  [mutationSetSelectedTheme](state: IThemeState, themeId: string) {
    state.selected = themeId;
  },
  [mutationSetUploadedFile](state: IThemeState, fileId: string) {
    state.uploadedFile = fileId;
  },
  [mutationSetThemedPdf](state: IThemeState, themedPdf: string) {
    state.themedPdf = themedPdf;
  },
  [mutationSetLoadingPdf](state: IThemeState, loading: boolean) {
    state.loadingPdfUpload = loading;
  },
  [mutationSetLoadedPdf](state: IThemeState, loaded: boolean) {
    state.loadedPdfUpload = loaded;
  },
  [mutationSetLoadingApplyTheme](state: IThemeState, loading: boolean) {
    state.loadingApplyTheme = loading;
  },
  [mutationSetLoadedApplyTheme](state: IThemeState, loaded: boolean) {
    state.loadedApplyTheme = loaded;
  }
};
