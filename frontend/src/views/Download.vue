<template>
  <div class="download">
    <h1>The deal has been scarified, you may take it...</h1>
    <a :href="url" target="_blank" class="download-monster button"
      >Fetch the monster</a
    >
  </div>
</template>

<script lang="ts">
import { Vue, Component, Prop } from "vue-property-decorator";
import { State, Action } from "vuex-class";
import { namespace as themeNamespace } from "@/store/themes/theme.store";

@Component
export default class Download extends Vue {
  @State("themedPdf", { namespace: themeNamespace }) url: string;

  mounted() {
    if (!this.url) {
      this.$router.go(-1);
    }
  }
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

https://scary-bucket.s3.eu-west-1.amazonaws.com/themed-documents/awesome-o.pdf?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAZD335QTAZ6FZMZU3%2F20191114%2Feu-west-1%2Fs3%2Faws4_request&X-Amz-Date=20191114T145835Z&X-Amz-Expires=900&X-Amz-SignedHeaders=host&X-Amz-Signature=2cf2830278f3f7142b85e9cb3c7f66b3858ddda2fd9812ad2978748bee14e948
