/*
 * Copyright (c) 2018 LynxiTech Inc - All rights reserved.
 * Author:曹凯
 * NOTICE: All information contained here is, and remains
 * the property of LynxiTech Incorporation. This file can not
 * be copied or distributed without permission of LynxiTech Inc.
 */

// 本文件实现了日志的中间件

// Package log provides log handling using logrus package.
//
// Based on github.com/stephenmuss/ginerus but adds more options.
package log

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// Ginrus returns a gin.HandlerFunc (middleware) that logs requests using logrus.
//
// Requests with errors are logged using logrus.Error().
// Requests without errors are logged using logrus.Info().
//
// It receives:
//   1. A time package format string (e.g. time.RFC3339).
//   2. A boolean stating whether to use UTC time zone or local.
func Logger() gin.HandlerFunc {
	timeFormat := time.RFC3339
	utc := true

	return func(c *gin.Context) {
		start := time.Now()
		// some evil middlewares modify this values
		path := c.Request.URL.Path
		c.Next()

		end := time.Now()
		latency := end.Sub(start)
		if utc {
			end = end.UTC()
		}
		status := c.Writer.Status()
		method := c.Request.Method
		time := end.Format(timeFormat)

		entry := log.WithFields(log.Fields{
			"module":     "api",
			"type":       "log",
			"status":     status,
			"method":     method,
			"path":       path,
			"ip":         c.ClientIP(),
			"latency":    latency,
			"user-agent": c.Request.UserAgent(),
			"time":       time,
		})

		if len(c.Errors) > 0 {
			// Append error field if this is an erroneous request.
			entry.Error(c.Errors.String())
		} else {
			ltms := float64(latency.Nanoseconds()) / 1000000.0
			msg := fmt.Sprintf("(%.3fms) %d - %s %s", ltms, status, method, path)
			if status >= 500 {
				entry.Error(msg)
			} else if status >= 400 {
				entry.Warn(msg)
			} else {
				entry.Info(msg)
			}
		}
	}
}
