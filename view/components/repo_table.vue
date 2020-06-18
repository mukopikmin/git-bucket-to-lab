<template>
  <div>
    <b-input-group prepend="Filter" class="mt-3">
      <b-form-input v-model="query"></b-form-input>
      <b-input-group-append>
        <b-button variant="outline-primary" @click="clear">
          <b-icon-x></b-icon-x>
        </b-button>
      </b-input-group-append>
    </b-input-group>

    <div class="card my-3">
      <table class="table card-table table-hover">
        <thead>
          <tr>
            <th></th>
            <th>GitBucket</th>
            <th>GitLab</th>
          </tr>
        </thead>
        <tbody>
          <tr v-if="loading">
            <td colspan="4">
              <div class="text-center my-3">
                <b-spinner variant="primary"></b-spinner>
              </div>
            </td>
          </tr>
          <tr v-for="(pair, i) in filtered" :key="i">
            <td>
              <nuxt-link
                :to="{
                  name: 'owner-name',
                  params: { owner: pair.repo.owner.login, name: pair.repo.name }
                }"
                >Migration</nuxt-link
              >
            </td>
            <td>
              <a target="_blank" :href="pair.repo.http_url">{{
                pair.repo.full_name
              }}</a>
            </td>
            <td v-if="pair.project">
              <a target="_blank" :href="pair.project.web_url">{{
                pair.project.path_with_namespace
              }}</a>
            </td>
            <td v-else></td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
import { mapState } from 'vuex'

export default {
  props: ['loading'],
  data() {
    return {
      query: ''
    }
  },
  computed: {
    ...mapState(['pairs']),
    filtered() {
      return this.pairs.filter((p) => p.repo.name.includes(this.query))
    }
  },
  methods: {
    clear() {
      this.query = ''
    }
  }
}
</script>

<style scoped>
table {
  margin: 0;
}
</style>
