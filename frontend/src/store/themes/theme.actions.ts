import { ActionTree } from "vuex";
import { RootState } from "..";
import { IThemeState } from "@/store/themes/theme.store";

export const actionLoadThemes = "loadTheme";

export const actions: ActionTree<IThemeState, RootState> = {
  async [actionLoadThemes]({ commit }) {}
};
