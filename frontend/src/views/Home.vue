<template>
  <div class="home">
    <Banner />
    <ThemeSelector @select="select" :selected="selected" :themes="themes" />
    <FileUpload @file="handleFile" /><br />
    <button v-if="uploadedFile" @click="handelDownload">
      Download your themed pdf
    </button>
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
import {
  actionLoadThemes,
  actionUploadFile
} from "../store/themes/theme.actions";

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
  @State("uploadedFile", { namespace: themeNamespace }) uploadedFile: string;

  @Action(actionSetTheme, { namespace: themeNamespace }) actionSetTheme: any;
  @Action(actionLoadThemes, { namespace: themeNamespace })
  actionLoadThemes: any;
  @Action(actionUploadFile, { namespace: themeNamespace })
  actionUploadFile: any;

  showWitch = false;

  mounted() {
    this.actionLoadThemes();
  }

  handleFile(file) {
    this.actionUploadFile(file);
  }

  handleDownload(downloadFile) {
    // Download from url
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
