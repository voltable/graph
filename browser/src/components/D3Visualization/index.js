import {Element as PolymerElement} from "@polymer/polymer/polymer-element.js";
import "@polymer/font-roboto/roboto.js";
import * as d3 from 'd3/build/d3'

import css from './style.scss';
import  template from './template.html';

export default class D3Visualization extends PolymerElement {

  static get template() {
    return `
    <style>     
      ${css}
    </style>
    ${template}`
  }

  constructor() {
    super();
    this.name = '3.0 preview';
  }

  ready() {
    super.ready();

    let canvas = this.$.canvas
    let context = canvas.getContext('2d')

    this.width = canvas.width
    this.height = canvas.height

    let color = d3.scaleOrdinal()
        .range(d3.schemeCategory20);

    let simulation = d3.forceSimulation()
        .force("charge", d3.forceManyBody().strength(-18))
        .force("link", d3.forceLink().iterations(4).id(function (d) { return d.id; }))
        .force("x", d3.forceX())
        .force("y", d3.forceY());

    d3.json("/json/graph.json", (error, graph) => {
        if (error) throw error;

        let users = d3.nest()
            .key(function (d) { return d.user; })
            .entries(graph.nodes)
            .sort(function (a, b) { return b.values.length - a.values.length; });

        color.domain(users.map(function (d) { return d.key; }));

        simulation
            .nodes(graph.nodes)
            .on("tick", () => {
              this._ticked.bind(this)(context, graph, users, color);
            });

        simulation.force("link")
            .links(graph.links);
    });
  }

  static get properties() {
    width: {
      Type: Number 
    }
    height: {
      Type: Number 
    }
  }

  _ticked(context, graph, users, color) {
    context.clearRect(0, 0, this.width, this.height);
    context.save();
    context.translate(this.width / 2, this.height / 2);

    context.beginPath();
    graph.links.forEach((d) => this._drawLink(context, d));
    context.strokeStyle = "#aaa";
    context.stroke();

    users.forEach((user) => {
        context.beginPath();
        user.values.forEach((d) => this._drawNode(context, d));
        context.fillStyle = color(user.key);
        context.fill();
    });

    context.restore();
  }

  _drawLink(context, d) {
    context.moveTo(d.source.x, d.source.y);
    context.lineTo(d.target.x, d.target.y);
  }

  _drawNode(context, d) {
    context.moveTo(d.x + 3, d.y);
    context.arc(d.x, d.y, 3, 0, 2 * Math.PI);
  }
}

customElements.define('cg-d3', D3Visualization);