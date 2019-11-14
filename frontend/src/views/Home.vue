<template>
  <div class="home">
    <Banner />
    <ThemeSelector @select="select" :selected="selected" :themes="themes" />
    <FileUpload />
  </div>
</template>

<script lang="ts">
import Banner from "@/components/Banner.vue";
import ThemeSelector from "@/components/ThemeSelector.vue";
import FileUpload from "@/components/FileUpload.vue";
import { Vue, Component } from "vue-property-decorator";
import { State, Action } from "vuex-class";
import { actionSetTheme } from "@/store/themes/theme.actions";
import { namespace as themeNamespace } from "@/store/themes/theme.store";
import { actionLoadThemes } from "../store/themes/theme.actions";

@Component({
  components: {
    Banner,
    ThemeSelector,
    FileUpload
  }
})
export default class Home extends Vue {
  @State("selected", { namespace: themeNamespace }) selected: string;
  @State("themes", { namespace: themeNamespace }) themes: [string];

  @Action(actionSetTheme, { namespace: themeNamespace }) actionSetTheme: any;
  @Action(actionLoadThemes, { namespace: themeNamespace })
  actionLoadThemes: any;

  showWitch = false;

  async mounted() {
    await this.actionLoadThemes();
  }
  select(index: number) {
    this.actionSetTheme(index);
  }
}
</script>
<style lang="scss" scoped>
.home {
}
</style>
