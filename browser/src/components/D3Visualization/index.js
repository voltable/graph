import { Element as PolymerElement } from "@polymer/polymer/polymer-element.js";
import "@polymer/font-roboto/roboto.js";
import * as d3 from 'd3/build/d3'

import css from './style.scss';
import template from './template.html';

export default class D3Visualization extends PolymerElement {

  static get template() {
    return `
    <style>     
      ${css}
    </style>
    ${template}`
  }

  connectedCallback() {
    super.connectedCallback();

    let canvas = this.$.canvas;

    let width = canvas.clientWidth;
    let height = canvas.clientHeight;

    let context = canvas.getContext('2d');

    canvas.width = width;
    canvas.height = height;

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
          this._ticked.bind(this)(context, graph, users, color, width, height);
        });

      simulation.force("link")
        .links(graph.links);
    });
  }

  _ticked(context, graph, users, color, width, height) {
    context.clearRect(0, 0, width, height);
    context.save();
    context.translate(width / 2, height / 2);

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