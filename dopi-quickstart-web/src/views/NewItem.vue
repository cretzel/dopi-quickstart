<template>

  <div class="new-item">
    <h1 class="title">Create Item</h1>
    <h2 class="subtitle">
      Enter the details of this item.
    </h2>


    <div v-if="item">

      <div class="columns">
        <div class="column is-half">

          <div class="field">
            <label class="label">Text</label>
            <div class="control">
              <input class="input" type="text" id="text" placeholder="Text" v-model="item.text" v-on:keyup.enter="saveItem"/>
            </div>
          </div>

        </div>
      </div>

    </div>

    <button v-on:click="saveItem" id="save" class="button is-primary mt-3">Save</button>

  </div>
</template>

<script>
import itemService from "../service/ItemService.js";
import store from "../store/Store.js";
import router from "@/router";

export default {
  name: "NewItem",
  data: function () {
    return {
      item: {
        text: "",
      },
    };
  },
  methods: {
    validate(itemDto) {
      const errors = [];
      if (!itemDto.text) {
        errors.push("Text is required")
      }
      return errors;
    },

    saveItem() {
      let itemDto = {
        text: this.item.text,
      };
           
      const errors = this.validate(itemDto);
      if (errors.length > 0) {
        store.setMessage({ texts: errors, type: "danger" });
        return
      }
      itemService
        .postItem(itemDto)
        .then(() => {
          router.push({ name: "ItemList" });
          store.setMessage({ text: "Item created", type: "success" });
        })
        .catch((e) => {
          console.log(e);
          store.setMessage({ text: e, type: "danger" });
        });
    },
  },
};
</script>
