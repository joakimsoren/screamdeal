<template>
  <div class="home">
    <Banner />
    <ThemeSelector />
  </div>
</template>

<script>
import Banner from "@/components/Banner.vue";
import ThemeSelector from "@/components/ThemeSelector.vue";
import { Vue, Component } from "vue-property-decorator";

@Component({
  components: {
    Banner,
    ThemeSelector
  }
})
export default class Home extends Vue {}
</script>
<style lang="scss" scoped>
.home {
}
</style>
