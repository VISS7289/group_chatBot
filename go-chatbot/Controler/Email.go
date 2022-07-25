package Controler

import (
	"github.com/gin-gonic/gin"
	"go-chatbot/Dao/Redis"
	"go-chatbot/Models"
	"go-chatbot/Pkg/Email"
	"go-chatbot/Pkg/MyRandom"
	"go.uber.org/zap"
)

func VerifiHandler(c *gin.Context) {
	var verifEmail Models.VerifiEmail
	if err := c.ShouldBindJSON(&verifEmail); err != nil {
		zap.L().Error("Verifi Parm Error", zap.Error(err))
		Models.ResponseError(c, Models.CodeInvalidParm)
		return
	}
	verifiCode := MyRandom.RandomString(6)
	if err := Email.SendEmail(verifEmail.To, "验证码", verifEmail.Username, verifiCode); err != nil {
		zap.L().Error("Verifi Send Error", zap.Error(err))
		Models.ResponseError(c, Models.CodeSendVerifiErr)
		return
	}
	if err := Redis.RedisSet(verifEmail.To, verifiCode); err != nil {
		zap.L().Error("Redis Set Error", zap.Error(err))
		Models.ResponseError(c, Models.CodeServerBusy)
		return
	}
	Models.ResponseSuccess(c, []string{})
}

func VerifiExam(c *gin.Context) {
	var verifExam Models.VerifiExam
	if err := c.ShouldBindJSON(&verifExam); err != nil {
		zap.L().Error("Verifi Parm Error", zap.Error(err))
		Models.ResponseError(c, Models.CodeInvalidParm)
		return
	}
	val, err := Redis.RedisFind(verifExam.Username)
	if err != nil {
		Models.ResponseError(c, Models.CodeVerifiNotFund)
		return
	}
	if val == verifExam.VerifiCode {
		Models.ResponseSuccess(c, []string{})
	} else {
		Models.ResponseError(c, Models.CodeVerifiErr)
	}
}
