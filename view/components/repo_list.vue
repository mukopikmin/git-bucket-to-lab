<template>
  <div>
    <b-row>
      <b-col>
        <b-input-group>
          <template v-slot:prepend>
            <b-input-group-text>
              <b-icon-person></b-icon-person>
            </b-input-group-text>
          </template>
          <b-form-select
            v-model="owner"
            :options="options"
            @change="onOwnerChange"
          ></b-form-select>
        </b-input-group>
      </b-col>
      <b-col>
        <b-input-group>
          <template v-slot:prepend>
            <b-input-group-text>
              <b-icon-search></b-icon-search>
            </b-input-group-text>
          </template>
          <b-form-input
            v-model="query"
            placeholder="Filter with repository name"
          ></b-form-input>
          <b-input-group-append>
            <b-button variant="outline-primary" @click="clear">
              <b-icon-x></b-icon-x>
            </b-button>
          </b-input-group-append>
        </b-input-group>
      </b-col>
    </b-row>

    <b-card-group columns class="mt-3">
      <b-card v-for="(pair, i) in pagedPairs" :key="i">
        <b-card-text>
          <nuxt-link
            :to="{
              name: 'owner-name',
              params: { owner: pair.repo.owner.login, name: pair.repo.name }
            }"
          >
            <b-icon-lock v-if="pair.repo.private" class="mr-1" />
            <b-icon-bookmark v-else class="mr-1" />
            {{ pair.repo.full_name }}
          </nuxt-link>
        </b-card-text>
        <b-card-text>
          <small>
            <a target="_blank" :href="pair.repo.http_url">
              <b-icon-box-arrow-up-right class="mr-1" />
              GitBucket
            </a>
            <a
              v-if="pair.project"
              target="_blank"
              :href="pair.project.web_url"
              class="ml-3"
            >
              <b-icon-box-arrow-up-right class="mr-1" />
              GitLab
            </a>
          </small>
        </b-card-text>
      </b-card>
    </b-card-group>

    <Pagination
      v-if="paginationEnabled"
      :page="page"
      :page-size="pageSize"
      @change="onPageChange"
    />

    <div v-if="loading" class="text-center my-5">
      <b-spinner variant="primary"></b-spinner>
    </div>
  </div>
</template>

<script>
import { mapState } from 'vuex'
import Pagination from '@/components/pagination'

export default {
  components: {
    Pagination
  },
  props: {
    loading: Boolean
  },
  data() {
    return {
      query: '',
      owner: null,
      page: 1,
      perPage: process.env.pageSize
    }
  },
  computed: {
    ...mapState(['pairs', 'gitbucketUser', 'gitbucketGroups']),
    filteredPairs() {
      return this.pairs
        .filter((p) => p.repo.name.includes(this.query))
        .filter((p) => !this.owner || p.repo.owner.login === this.owner)
    },
    pagedPairs() {
      return this.filteredPairs.slice(
        this.perPage * (this.page - 1),
        this.perPage * this.page
      )
    },
    pageSize() {
      return Math.ceil(this.filteredPairs.length / this.perPage)
    },
    paginationEnabled() {
      return this.pageSize > 1
    },
    options() {
      if (!this.gitbucketUser) {
        return []
      }

      const none = {
        value: null,
        text: 'Select user / group to show'
      }
      const user = {
        value: this.gitbucketUser.login,
        text: this.gitbucketUser.login
      }
      const groups = this.gitbucketGroups.map((g) => {
        return {
          value: g.login,
          text: g.login
        }
      })

      return [none, user, ...groups]
    }
  },
  mounted() {
    this.perPage = (this.perPage + (3 - (this.perPage % 3))) * 2
  },
  methods: {
    clear() {
      this.query = ''
    },
    onPageChange(e) {
      this.page = e
    },
    onOwnerChange() {
      this.page = 1
    }
  }
}
</script>

<style scoped>
table {
  margin: 0;
}
</style>
