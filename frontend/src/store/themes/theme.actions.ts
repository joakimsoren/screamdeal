import { ActionTree } from "vuex";
import { RootState } from "..";
import { IThemeState } from "@/store/themes/theme.store";
import { mutationSetSelectedTheme, mutationSetThemes } from "./theme.mutations";

export const actionLoadThemes = "loadTheme";
export const actionSetTheme = "setTheme";

export const actions: ActionTree<IThemeState, RootState> = {
  async [actionLoadThemes]({ commit }) {
    const response = await fetch("http://localhost:8080/themes");
    const data = await response.json();
    commit(mutationSetThemes, data.urls);
  },
  async [actionSetTheme]({ commit }, themeId: string) {
    commit(mutationSetSelectedTheme, themeId);
  }
};
