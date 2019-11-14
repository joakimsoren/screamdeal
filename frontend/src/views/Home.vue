<template>
  <div class="home">
    <Banner />
    <ThemeSelector />
    <FileUpload />
  </div>
</template>

<script>
import Banner from "@/components/Banner.vue";
import ThemeSelector from "@/components/ThemeSelector.vue";
import FileUpload from "@/components/FileUpload.vue";
import { Vue, Component } from "vue-property-decorator";

@Component({
  components: {
    Banner,
    ThemeSelector,
    FileUpload
  }
})
export default class Home extends Vue {}
</script>
<style lang="scss" scoped>
.home {
}
</style>
