import { ActionTree } from "vuex";
import { RootState } from "..";
import { IThemeState } from "@/store/themes/theme.store";
import { mutationSetSelectedTheme } from "./theme.mutations";

export const actionLoadThemes = "loadTheme";
export const actionSetTheme = "setTheme";

export const actions: ActionTree<IThemeState, RootState> = {
  async [actionLoadThemes]({ commit }) {},
  async [actionSetTheme]({ commit }, themeId: number) {
    commit(mutationSetSelectedTheme, themeId);
  },
};
