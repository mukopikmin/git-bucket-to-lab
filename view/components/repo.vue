<template>
  <b-card no-body>
    <b-card-header header-tag="nav">
      <span class="title">Repository</span>
      <MigrateButton class="migrate-button" :action="migrateRepo" />
    </b-card-header>

    <b-card-body>
      <div v-if="loading" class="text-center my-2">
        <b-spinner variant="primary"></b-spinner>
      </div>

      <div v-else-if="repo">
        <b-card-title>{{ repo.full_name }}</b-card-title>
        <b-card-text>{{ repo.description }}</b-card-text>
      </div>
    </b-card-body>

    <div v-if="repo">
      <b-card-body v-if="isNoBranches">
        <b-card-text>No branches</b-card-text>
      </b-card-body>

      <div v-else>
        <b-list-group flush>
          <b-list-group-item>Branches</b-list-group-item>
          <b-list-group-item
            v-for="branch in repo.branches"
            :key="branch.commit.sha"
            class="d-flex justify-content-between align-items-center text-align-left"
          >
            <Branch
              class="branch"
              :name="branch.name"
              :sha="branch.commit.sha"
            />
          </b-list-group-item>
        </b-list-group>
      </div>
    </div>
  </b-card>
</template>

<script>
import { mapActions, mapState } from 'vuex'
import Branch from '@/components/branch'
import MigrateButton from '@/components/migrate_button'

export default {
  components: {
    Branch,
    MigrateButton
  },
  props: ['repo', 'loading'],
  computed: {
    isNoBranches() {
      return this.repo.branches.length === 0
    },
    ...mapState(['gitbucketUser', 'gitbucketToken', 'gitlabToken'])
  },
  methods: {
    ...mapActions(['setRepo', 'setProject', 'setError']),
    async migrateRepo() {
      try {
        const res = await this.$axios.$post(
          `/${this.repo.owner.login}/${this.repo.name}/repo`,
          null,
          {
            headers: {
              'X-GITBUCKET-USER': this.gitbucketUser.login,
              'X-GITBUCKET-TOKEN': this.gitbucketToken,
              'X-GITLAB-TOKEN': this.gitlabToken
            }
          }
        )

        this.setRepo(res.repo)
        this.setProject(res.project)
        this.setError(null)
      } catch (e) {
        this.setError(e.response.data.message)
      }
    }
  }
}
</script>

<style scoped>
.card-header {
  padding-top: 8.5px;
  padding-bottom: 8.5px;
}
.title {
  height: 100%;
  vertical-align: middle;
}
.migrate-button {
  float: right;
}
.branch {
  width: 100%;
}
</style>
