<template>
  <b-card no-body>
    <b-card-header header-tag="nav">
      <span class="title">Issues</span>
      <MigrateButton :action="migrateIssues" />
    </b-card-header>

    <div v-if="loading" class="text-center my-2">
      <b-spinner variant="primary"></b-spinner>
    </div>

    <div v-else>
      <b-list-group flush>
        <b-list-group-item
          v-for="issue in issues"
          :key="`gitbcket-issue-${issue.number}`"
          class="d-flex justify-content-between align-items-center text-align-left"
        >
          <span>
            <b-icon-check-circle
              v-if="issue.state == 'open'"
              variant="success"
              class="mr-1"
            ></b-icon-check-circle>
            <b-icon-slash-circle
              v-else
              variant="danger"
              class="mr-1"
            ></b-icon-slash-circle>
            #{{ issue.number }} {{ issue.title }}
          </span>
          <b-badge v-if="issue.comments.length > 0" variant="primary" pill>{{
            issue.comments.length
          }}</b-badge>
        </b-list-group-item>
      </b-list-group>
    </div>
  </b-card>
</template>

<script>
import { mapActions, mapState } from 'vuex'
import MigrateButton from '@/components/migrate_button'

export default {
  components: {
    MigrateButton
  },
  props: ['repo', 'issues', 'loading'],
  computed: {
    ...mapState(['gitbucketToken', 'gitlabToken'])
  },
  methods: {
    ...mapActions(['setRepo', 'setProject']),
    async migrateIssues() {
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

      this.setRepo(res.repo)
      this.setProject(res.project)
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
