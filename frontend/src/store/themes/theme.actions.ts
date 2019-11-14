import { ActionTree } from "vuex";
import { RootState } from "..";
import { IThemeState } from "@/store/themes/theme.store";
import {
  mutationSetSelectedTheme,
  mutationSetThemes,
  mutationSetUploadedFile,
  mutationSetThemedPdf,
  mutationSetLoadingPdf,
  mutationSetLoadedPdf,
  mutationSetLoadingApplyTheme,
  mutationSetLoadedApplyTheme
} from "./theme.mutations";
import router from "@/router";

export const actionLoadThemes = "loadTheme";
export const actionSetTheme = "setTheme";
export const actionUploadFile = "uploadFile";
export const actionApplyThemeToPdf: string = "applyThemeToPdf";

export const actions: ActionTree<IThemeState, RootState> = {
  async [actionLoadThemes]({ commit }) {
    const response = await fetch("http://localhost:8080/themes");
    const data = await response.json();
    commit(mutationSetThemes, data.urls);
  },
  async [actionSetTheme]({ commit }, themeId: string) {
    commit(mutationSetSelectedTheme, themeId);
  },
  async [actionUploadFile]({ commit }, file: File) {
    commit(mutationSetLoadingPdf, true);
    try {
      const response = await fetch("http://localhost:8080/put-pdf", {
        method: "POST",
        body: file
      });
      const data = await response.json();
      commit(mutationSetUploadedFile, data);
      commit(mutationSetLoadedPdf, true);
    } catch {}
    commit(mutationSetLoadingPdf, false);
  },
  async [actionApplyThemeToPdf](
    { commit },
    uploadedFile: string,
    theme: string
  ) {
    commit(mutationSetLoadingApplyTheme, true);
    try {
      const response = await fetch("http://localhost:8080/add-theme-to-pdf", {
        method: "POST",
        body: JSON.stringify({
          file: uploadedFile,
          theme
        }),
        headers: {
          "Content-Type": "application/json"
        }
      });
      const data = await response.json();
      commit(mutationSetThemedPdf, data.pdf);
      router.push("/download");
    } catch {}
    commit(mutationSetLoadingApplyTheme, false);
    commit(mutationSetLoadedApplyTheme, true);
  }
};
