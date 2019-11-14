import { Module } from "vuex";
import { RootState } from "..";
import { actions } from "./theme.actions";
import { mutations } from "./theme.mutations";

export const namespaced: boolean = true;
export const namespace: string = "themes";

export interface IThemeState {
  loadingPdfUpload: boolean;
  loadedPdfUpload: boolean;
  loadingApplyTheme: boolean;
  loadedApplyTheme: boolean;
  themes: string[];
  selected?: string;
  uploadedFile?: string;
  themedPdf?: string;
}

export const state: IThemeState = {
  loadingPdfUpload: false,
  loadedPdfUpload: false,
  loadingApplyTheme: false,
  loadedApplyTheme: false,
  themes: [],
  selected: undefined,
  uploadedFile: undefined,
  themedPdf: undefined
};

export const themes: Module<IThemeState, RootState> = {
  state,
  actions,
  mutations,
  namespaced
};
