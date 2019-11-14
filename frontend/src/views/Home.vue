<template>
  <div class="home">
    <Banner />
    <ThemeSelector @select="select" :selected="selected" />
  </div>
</template>

<script lang="ts">
import Banner from "@/components/Banner.vue";
import ThemeSelector from "@/components/ThemeSelector.vue";
import { Vue, Component } from "vue-property-decorator";
import { State, Action } from "vuex-class";
import { actionSetTheme } from "@/store/themes/theme.actions";
import { namespace as themeNamespace } from "@/store/themes/theme.store";

@Component({
  components: {
    Banner,
    ThemeSelector
  }
})
export default class Home extends Vue {
  @State("selected", { namespace: themeNamespace }) selected: number;
  @Action(actionSetTheme, { namespace: themeNamespace }) actionSetTheme: any;

  select(index: number) {
    this.actionSetTheme(index);
  }
}
</script>
<style lang="scss" scoped>
.home {
}
</style>
