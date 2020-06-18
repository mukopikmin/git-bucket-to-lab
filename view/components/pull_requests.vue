<template>
  <b-card no-body>
    <b-card-header header-tag="nav">
      <span class="title">Pull Requests</span>
      <MigrateButton class="migrate-button" :action="migratePulls" />
    </b-card-header>

    <b-card-body v-if="loading" class="text-center my-2">
      <b-spinner variant="primary"></b-spinner>
    </b-card-body>

    <div v-else>
      <b-list-group flush>
        <b-list-group-item v-if="noPulls">No pull requests</b-list-group-item>
        <b-list-group-item
          v-for="pull in pagedPulls"
          :key="`pull-${pull.number}`"
          class="d-flex justify-content-between align-items-center text-align-left"
        >
          <span>
            <b-icon-check-circle
              v-if="pull.state == 'open'"
              variant="success"
              class="mr-1"
            ></b-icon-check-circle>
            <b-icon-slash-circle
              v-else
              variant="danger"
              class="mr-1"
            ></b-icon-slash-circle>
            #{{ pull.number }} {{ pull.title }}
          </span>
          <b-badge v-if="pull.comments.length > 0" variant="primary" pill>{{
            pull.comments.length
          }}</b-badge>
        </b-list-group-item>
      </b-list-group>

      <Pagination
        v-if="paginationEnabled"
        :page="page"
        :page-size="pageSize"
        @change="onPageChange"
      />
    </div>
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
  props: ['repo', 'pulls', 'loading'],
  data() {
    return {
      page: 1,
      perPage: 4
    }
  },
  computed: {
    ...mapState(['gitbucketToken', 'gitlabToken']),
    pagedPulls() {
      return this.pulls.slice(
        this.perPage * (this.page - 1),
        this.perPage * this.page
      )
    },
    pageSize() {
      return Math.ceil(this.pulls.length / this.perPage)
    },
    paginationEnabled() {
      return this.pageSize > 1
    },
    noPulls() {
      return this.pulls.length === 0
    }
  },
  methods: {
    ...mapActions(['setRepo', 'setProject', 'setError']),
    onPageChange(e) {
      this.page = e
    },
    async migratePulls() {
      try {
        const res = await this.$axios.$post(
          `/${this.repo.owner.login}/${this.repo.name}/pulls`,
          null,
          {
            headers: {
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
</style>
