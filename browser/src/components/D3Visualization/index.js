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
    console.log(canvas);
    let context = canvas.getContext('2d')

    let width = canvas.width
    let height = canvas.height

    var color = d3.scaleOrdinal()
        .range(d3.schemeCategory20);

    var simulation = d3.forceSimulation()
        .force("charge", d3.forceManyBody().strength(-18))
        .force("link", d3.forceLink().iterations(4).id(function (d) { return d.id; }))
        .force("x", d3.forceX())
        .force("y", d3.forceY());

    d3.json("/json/graph.json", function (error, graph) {
        if (error) throw error;

        var users = d3.nest()
            .key(function (d) { return d.user; })
            .entries(graph.nodes)
            .sort(function (a, b) { return b.values.length - a.values.length; });

        color.domain(users.map(function (d) { return d.key; }));

        simulation
            .nodes(graph.nodes)
            .on("tick", ticked);

        simulation.force("link")
            .links(graph.links);

        function ticked() {
            context.clearRect(0, 0, width, height);
            context.save();
            context.translate(width / 2, height / 2);

            context.beginPath();
            graph.links.forEach(drawLink);
            context.strokeStyle = "#aaa";
            context.stroke();

            users.forEach(function (user) {
                context.beginPath();
                user.values.forEach(drawNode);
                context.fillStyle = color(user.key);
                context.fill();
            });

            context.restore();
        }
    });

    function drawLink(d) {
        context.moveTo(d.source.x, d.source.y);
        context.lineTo(d.target.x, d.target.y);
    }
    
    function drawNode(d) {
        context.moveTo(d.x + 3, d.y);
        context.arc(d.x, d.y, 3, 0, 2 * Math.PI);
    }
  }

  static get properties() {
    name: {
      Type: String
    }
  }
}

customElements.define('cg-d3', D3Visualization);