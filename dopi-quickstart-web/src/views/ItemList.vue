<template>
  <div class="item-list">
    <h1 class="title">Items</h1>
    <h2 class="subtitle">
      View all your dopi items, create new items.
    </h2>

    <div class="columns is-pulled-right">
      <div class="column">
        <router-link
          id="create-item-button"
          :to="{ name: 'NewItem' }"
          tag="button"
          class="button is-primary mt-3"
          >New Item</router-link
        >
      </div>
    </div>

    <table class="table is-hoverable is-fullwidth">
      <thead>
        <tr>
          <th>Text</th>
          <th>Created</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="item in items" :key="item.id" v-bind:data-item="item.id">
          <td>
            {{ item.text }}
          </td>
          <td class="has-text-grey">
            {{ item.createdAt | formatDate }}
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
import itemService from "../service/ItemService.js";
import store from "../store/Store.js";

export default {
  name: "ItemList",
  data: function () {
    return {
      items: [],
    };
  },
  methods: {
    getItems() {
      itemService
        .getItems()
        .then((itemDtos) => {
          this.items = itemDtos;
        })
        .catch(() => {
          store.setMessage({ text: "Cannot get items", type: "danger" });
        });
    },
  },
  created: function () {
    this.getItems();
  },
};
</script>
