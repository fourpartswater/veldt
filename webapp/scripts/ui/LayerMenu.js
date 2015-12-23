(function() {
    'use strict';

    var $ = require('jquery');

    var LayerMenu = function(spec) {
        var self = this;
        // parse inputs
        spec = spec || {};
        spec.label = spec.label !== undefined ? spec.label : 'missing-layer-name';
        spec.layer = spec.layer || null;
        spec.isVisible = !spec.layer.isHidden();
        spec.isMinimized = false;
        if (!spec.layer) {
            console.error('LayerMenu constructor must be passed a "layer" attribute');
            return;
        }
        // set layer
        this._layer = spec.layer;
        // set minimized
        this._minimized = spec.isMinimized;
        // create elements
        this._$menu = $(Templates.layermenu(spec));
        this._$head = this._$menu.find('.layer-control-head');
        this._$body = this._$menu.find('.layer-control-body');
        this._$title = this._$menu.find('.layer-title');
        this._$toggleIcon = this._$menu.find('.layer-toggle i');
        this._$toggle = this._$menu.find('.layer-toggle');
        this._$minimizeIcon = this._$menu.find('.minimize-toggle i');
        this._$minimize = this._$menu.find('.minimize-toggle');
        // attach callbacks
        this._$toggle.click(function() {
            self.toggleEnabled();
        });
        this._$title.click(function() {
            self.toggleMinimized();
        });
    };

    LayerMenu.prototype.toggleMinimized = function() {
        if (this._minimized) {
            this.maximize();
        } else {
            this.minimize();
        }
    };

    LayerMenu.prototype.minimize = function() {
        this._originalHeight = this._$body.outerHeight();
        this._$minimizeIcon.removeClass('fa-minus');
        this._$minimizeIcon.addClass('fa-plus');
        this._$body.css({
            height: '0px',
            border: 'none'
        });
        this._$head.css({
            padding: 'initial'
        });
        this._minimized = true;
    };

    LayerMenu.prototype.maximize = function() {
        this._$minimizeIcon.removeClass('fa-plus');
        this._$minimizeIcon.addClass('fa-minus');
        this._$body.css({
            height: this._originalHeight,
            border: ''
        });
        this._$head.css({
            padding: ''
        });
        this._minimized = false;
    };

    LayerMenu.prototype.toggleEnabled = function() {
        if (this._layer.isHidden()) {
            this.enable();
        } else {
            this.disable();
        }
    };

    LayerMenu.prototype.disable = function() {
        this._layer.hide();
        this._$toggleIcon.removeClass('fa-check-square-o');
        this._$toggleIcon.addClass('fa-square-o');
        this._$body.addClass('disabled');
    };

    LayerMenu.prototype.enable = function() {
        this._layer.show();
        this._$toggleIcon.removeClass('fa-square-o');
        this._$toggleIcon.addClass('fa-check-square-o');
        this._$body.removeClass('disabled');
    };

    LayerMenu.prototype.getElement = function() {
        return this._$menu;
    };

    LayerMenu.prototype.getBody = function() {
        return this._$body;
    };

    module.exports = LayerMenu;

}());
