<template>
  <b-card no-body>
    <b-card-header header-tag="nav">
      <span class="title">Repository</span>
      <MigrateButton
        class="migrate-button"
        :migrating="migrating"
        :action="migrateRepo"
      />
    </b-card-header>

    <b-card-body>
      <div v-if="loading" class="text-center my-2">
        <b-spinner variant="primary"></b-spinner>
      </div>

      <div v-else-if="repo">
        <b-card-title>
          <b-icon-lock v-if="repo.private" class="mr-1" />
          <b-icon-bookmarks v-else class="mr-1" />
          {{ repo.full_name }}
        </b-card-title>
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

    <template v-if="repo" v-slot:footer>
      <small class="text-muted">
        <a :href="repo.html_url" target="_blank">
          <b-icon-box-arrow-up-right class="mr-1" />
          Open repository
        </a>
      </small>
    </template>
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
  props: {
    repo: {
      type: Object,
      required: false,
      default: null
    },
    loading: Boolean
  },
  data() {
    return {
      migrating: false
    }
  },
  computed: {
    isNoBranches() {
      return this.repo.branches.length === 0
    },
    ...mapState(['username', 'gitbucketUser', 'gitbucketToken', 'gitlabToken'])
  },
  methods: {
    ...mapActions(['setRepo', 'setProject', 'setError']),
    async migrateRepo() {
      try {
        this.migrating = true
        const res = await this.$axios.$post(
          `/${this.repo.owner.login}/${this.repo.name}/repo`,
          null,
          {
            headers: {
              'X-GITBUCKET-USER': this.username,
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
      } finally {
        this.migrating = false
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
.divider {
  margin: 0;
}
</style>
