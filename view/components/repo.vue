<template>
  <b-card no-body>
    <b-card-header header-tag="nav">
      <span class="title">Repository</span>
      <MigrateButton
        class="migrate-button"
        :migrating="migrating"
        :action="migrateRepo"
        :label="buttonLabel"
        :disable="isNotMigratable"
      />
    </b-card-header>

    <b-card-body>
      <div v-if="loading" class="text-center my-2">
        <b-spinner variant="primary"></b-spinner>
      </div>

      <div v-else-if="repo">
        <b-card-title>
          <b-icon-lock v-if="repo.private" class="mr-1" />
          <b-icon-bookmark v-else class="mr-1" />
          {{ repo.owner.login }} / {{ repo.name }}
        </b-card-title>
        <b-card-text>{{ repo.description }}</b-card-text>
      </div>
    </b-card-body>

    <div v-if="!loading && repo">
      <h5 class="list-title">Branches</h5>
      <b-list-group flush class="border-top border-bottom">
        <b-list-group-item v-if="isNoBranches">No branches</b-list-group-item>
        <b-list-group-item
          v-for="branch in repo.branches"
          :key="branch.commit.sha"
          class="d-flex justify-content-between align-items-center text-align-left"
        >
          <Branch class="branch" :name="branch.name" :sha="branch.commit.sha" />
        </b-list-group-item>
      </b-list-group>

      <h5 class="list-title mt-4">Tags</h5>
      <b-list-group flush class="border-top">
        <b-list-group-item v-if="isNoTags">No tags</b-list-group-item>
        <b-list-group-item
          v-for="tag in repo.tags"
          :key="tag.name"
          class="d-flex justify-content-between align-items-center text-align-left"
        >
          <Branch class="branch" :name="tag.name" :sha="tag.sha" />
        </b-list-group-item>
      </b-list-group>
    </div>

    <template v-if="!loading && repo" v-slot:footer>
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
      migrating: false,
      repoMigrablae: false
    }
  },
  computed: {
    ...mapState([
      'username',
      'gitbucketUser',
      'gitbucketToken',
      'gitlabToken',
      'migratable',
      'project'
    ]),
    isNoBranches() {
      return this.repo.branches.length === 0
    },
    isNoTags() {
      return this.repo.tags.length === 0
    },
    isOrgRepo() {
      return this.repo.owner.type === 'Organization'
    },
    buttonLabel() {
      return this.project ? 'Sync' : 'Migrate'
    },
    isNotMigratable() {
      return this.loading || !this.migratable.repo
    }
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
              'X-GITBUCKET-TOKEN': this.gitbucketToken,
              'X-GITLAB-TOKEN': this.gitlabToken
            }
          }
        )

        this.setRepo({
          repo: res.repo,
          repoMigratable: res.repo_migratable,
          issuesMigratable: res.issues_migratable,
          pullsMigratable: res.pulls_migratable
        })
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
.list-title {
  padding-left: 20px;
}
.border-top {
  border-top: 1px solid rgba(0, 0, 0, 0.125);
}
.border-bottom {
  border-bottom: 1px solid rgba(0, 0, 0, 0.125);
}
</style>
