<template>
  <div>
    <hr class="divider" />
    <b-pagination-nav
      v-model="innerPage"
      :value="innerPage"
      :number-of-pages="pageSize"
      :link-gen="linkGen"
      align="center"
    ></b-pagination-nav>
  </div>
</template>

<script>
export default {
  props: {
    page: {
      type: Number,
      required: true
    },
    pageSize: {
      type: Number,
      required: true
    }
  },
  data() {
    return {
      perPage: 4
    }
  },
  computed: {
    innerPage: {
      get() {
        return this.page
      },
      set(value) {
        this.$emit('change', value)
      }
    }
  },
  methods: {
    linkGen(e) {
      return ''
    }
  }
}
</script>

<style scoped>
.divider {
  margin-top: 0;
}
</style>
