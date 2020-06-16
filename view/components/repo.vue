<template>
  <div class="card">
    <div class="card-header">Repository</div>
    <div class="card-body">
      <div v-if="loading" class="text-center my-2">
        <b-spinner variant="primary"></b-spinner>
      </div>

      <div v-else-if="repo">
        <h4 class="card-title">{{ repo.full_name }}</h4>
        <p>{{ repo.description }}</p>

        <p v-if="isNoBranches">No branches</p>
        <div v-else>
          <h5>Branches</h5>
          <ul>
            <li v-for="branch in repo.branches" :key="branch.commit.sha">
              {{ branch.name }} ({{ branch.commit.sha }})
            </li>
          </ul>
        </div>

        <button class="btn btn-outline-primary" @click="migrateRepo">
          Migrate
        </button>
      </div>
    </div>
  </div>
</template>

<script>
import { mapState } from 'vuex'

export default {
  props: ['repo', 'loading'],
  computed: {
    isNoBranches() {
      return this.repo.branches.length === 0
    },
    ...mapState(['gitbucketToken', 'gitlabToken'])
  },
  methods: {
    async migrateRepo() {
      const res = await this.$axios.$post(
        `/${this.repo.owner.login}/${this.repo.name}/repo`,
        null,
        {
          headers: {
            'X-GITBUCKET-TOKEN': this.gitbucketToken,
            'X-GITLAB-TOKEN': this.gitlabToken
          }
        }
      )

      this.project = res.Project
    }
  }
}
</script>

<style></style>
