import {Element as PolymerElement} from "@polymer/polymer/polymer-element.js";
import "@polymer/font-roboto/roboto.js";
import "@polymer/iron-icons/iron-icons.js";
import "@polymer/paper-icon-button/paper-icon-button.js";
import "@polymer/paper-item/paper-icon-item.js";
import "@polymer/app-layout/app-drawer-layout/app-drawer-layout.js";
import "@polymer/app-layout/app-drawer/app-drawer.js";
import "@polymer/app-layout/app-header/app-header.js";
import "@polymer/app-layout/app-header-layout/app-header-layout.js";
import "@polymer/app-layout/app-toolbar/app-toolbar.js";
import "@polymer/app-layout/app-scroll-effects/app-scroll-effects.js";
import "@polymer/iron-icon/iron-icon.js";

import '../D3Visualization';


import css from './style.scss';
import template from './template.html';

export default class App extends PolymerElement {

  static get template() {
    return `
    <style>     
      ${css}
    </style>
    ${template}`
  }

  constructor() {
    super();
  }

  static get properties() {
  }
}

customElements.define('cg-app', App);