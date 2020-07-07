<template>
  <b-card no-body>
    <b-card-header header-tag="nav">
      <span class="title">Issues</span>
      <MigrateButton
        class="migrate-button"
        :migrating="migrating"
        :action="migrateIssues"
        label="Migrate"
        :disable="isNotMigratable"
      />
    </b-card-header>

    <b-card-body v-if="loading" class="text-center my-2">
      <b-spinner variant="primary"></b-spinner>
    </b-card-body>

    <div v-else>
      <b-list-group flush :class="paginationEnabled ? 'border-bottom' : ''">
        <b-list-group-item v-if="noIssues">No issuess</b-list-group-item>
        <b-list-group-item
          v-for="issue in pagedIssues"
          :key="`gitbcket-issue-${issue.number}`"
          class="d-flex justify-content-between align-items-center text-align-left"
        >
          <span>
            <b-icon-info-circle
              v-if="issue.state === 'open'"
              variant="success"
              class="mr-1"
            />
            <b-icon-check2 v-else variant="danger" class="mr-1" />
            #{{ issue.number }} {{ issue.title }}
          </span>
          <b-badge v-if="issue.comments.length > 0" variant="primary" pill>
            {{ issue.comments.length }}
          </b-badge>
        </b-list-group-item>
      </b-list-group>

      <Pagination
        v-if="paginationEnabled"
        :page="page"
        :page-size="pageSize"
        @change="onPageChange"
      />
    </div>

    <template v-slot:footer>
      <small class="text-muted">
        <b-icon-info-circle class="mr-1" />
        {{ openCount }} Open
        <b-icon-check2 class="ml-2 mr-1" />
        {{ closedCount }} Closed
      </small>
    </template>
  </b-card>
</template>

<script>
import { mapActions, mapState } from 'vuex'
import MigrateButton from '@/components/migrate_button'
import Pagination from '@/components/pagination'

export default {
  components: {
    MigrateButton,
    Pagination
  },
  props: {
    repo: {
      type: Object,
      required: false,
      default: null
    },
    issues: {
      type: Array,
      required: true
    },
    loading: Boolean
  },
  data() {
    return {
      page: 1,
      perPage: process.env.pageSize,
      migrating: false
    }
  },
  computed: {
    ...mapState(['gitbucketToken', 'gitlabToken', 'migratable']),
    pagedIssues() {
      return this.issues.slice(
        this.perPage * (this.page - 1),
        this.perPage * this.page
      )
    },
    pageSize() {
      return Math.ceil(this.issues.length / this.perPage)
    },
    paginationEnabled() {
      return this.pageSize > 1
    },
    noIssues() {
      return this.issues.length === 0
    },
    openCount() {
      return this.issues.filter((i) => i.state === 'open').length
    },
    closedCount() {
      return this.issues.filter((i) => i.state === 'closed').length
    },
    isNotMigratable() {
      return this.loading || !this.migratable.issues
    }
  },
  methods: {
    ...mapActions(['setRepo', 'setProject', 'setError']),
    onPageChange(e) {
      this.page = e
    },
    async migrateIssues() {
      try {
        this.migrating = true
        const res = await this.$axios.$post(
          `/${this.repo.owner.login}/${this.repo.name}/issues`,
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
.border-bottom {
  border-bottom: 1px solid rgba(0, 0, 0, 0.125);
  margin-bottom: 16px;
}
</style>
