<template>
  <div class="download">
    <h1>The deal has been scarified, you may take it...</h1>
    <button class="download-monster">Fetch the monster</button>
  </div>
</template>

<script lang="ts">
import { Vue, Component, Prop } from "vue-property-decorator";

@Component
export default class Download extends Vue {
  @Prop() url: string;
}
</script>
<style lang="scss" scoped>
.download {
  color: white;

  h1 {
    font-size: 4rem;
  }
}
</style>
