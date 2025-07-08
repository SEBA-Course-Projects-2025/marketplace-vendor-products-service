package amqp

import (
	"dev-vendor/product-service/internal/shared/utils"
	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
	"os"
)

type AMQPConfig struct {
	Channel             *amqp.Channel
	ConfirmationChannel <-chan amqp.Confirmation
}

func SetUpExchange(channel *amqp.Channel) error {

	if err := channel.ExchangeDeclare("product.catalog.events", "direct", true, false, false, false, nil); err != nil {
		return utils.ErrorHandler(err, err.Error())
	}

	if _, err := channel.QueueDeclare("product.created.catalog", true, false, false, false, nil); err != nil {
		return utils.ErrorHandler(err, err.Error())
	}

	if err := channel.QueueBind("product.created.catalog", "product.created.catalog", "product.catalog.events", false, nil); err != nil {
		return utils.ErrorHandler(err, err.Error())
	}

	if _, err := channel.QueueDeclare("product.updated.catalog", true, false, false, false, nil); err != nil {
		return utils.ErrorHandler(err, err.Error())
	}

	if err := channel.QueueBind("product.updated.catalog", "product.updated.catalog", "product.catalog.events", false, nil); err != nil {
		return utils.ErrorHandler(err, err.Error())
	}

	if err := channel.ExchangeDeclare("vendor.product.events", "direct", true, false, false, false, nil); err != nil {
		return utils.ErrorHandler(err, err.Error())
	}
	
	if err := channel.QueueBind("vendor.product.quantity.checked", "vendor.product.quantity.checked", "vendor.product.events", false, nil); err != nil {
		return utils.ErrorHandler(err, err.Error())
	}

	if _, err := channel.QueueDeclare("vendor.check.product.quantity", true, false, false, false, amqp.Table{"x-dead-letter-exchange": "vendor.product.dlx", "x-dead-letter-routing-key": "vendor.check.product.quantity.dlq"}); err != nil {
		return utils.ErrorHandler(err, err.Error())
	}

	if err := channel.QueueBind("vendor.check.product.quantity", "vendor.check.product.quantity", "vendor.product.events", false, nil); err != nil {
		return utils.ErrorHandler(err, err.Error())
	}

	if _, err := channel.QueueDeclare("vendor.cancel.product.order", true, false, false, false, amqp.Table{"x-dead-letter-exchange": "vendor.product.dlx", "x-dead-letter-routing-key": "vendor.cancel.product.order.dlq"}); err != nil {
		return utils.ErrorHandler(err, err.Error())
	}

	if err := channel.QueueBind("vendor.cancel.product.order", "vendor.cancel.product.order", "vendor.product.events", false, nil); err != nil {
		return utils.ErrorHandler(err, err.Error())
	}

	return nil

}

func SetUpDlq(channel *amqp.Channel) error {

	if err := channel.ExchangeDeclare("vendor.product.dlx", "direct", true, false, false, false, nil); err != nil {
		return utils.ErrorHandler(err, err.Error())
	}

	if _, err := channel.QueueDeclare("vendor.check.product.quantity.dlq", true, false, false, false, amqp.Table{"x-dead-letter-exchange": "vendor.product.dlx", "x-dead-letter-routing-key": "vendor.check.product.quantity.dlq"}); err != nil {
		return utils.ErrorHandler(err, err.Error())
	}

	if err := channel.QueueBind("vendor.check.product.quantity.dlq", "vendor.check.product.quantity.dlq", "vendor.product.dlx", false, nil); err != nil {
		return utils.ErrorHandler(err, err.Error())
	}

	if _, err := channel.QueueDeclare("vendor.cancel.product.order.dlq", true, false, false, false, amqp.Table{"x-dead-letter-exchange": "vendor.product.dlx", "x-dead-letter-routing-key": "vendor.cancel.product.order.dlq"}); err != nil {
		return utils.ErrorHandler(err, err.Error())
	}

	if err := channel.QueueBind("vendor.cancel.product.order.dlq", "vendor.cancel.product.order.dlq", "vendor.product.dlx", false, nil); err != nil {
		return utils.ErrorHandler(err, err.Error())
	}

	return nil

}

func ConnectAMQP() (*AMQPConfig, error) {

	if err := godotenv.Load(); err != nil {
		return nil, utils.ErrorHandler(err, err.Error())
	}

	amqpUrl := os.Getenv("AMQP_URL")

	connection, err := amqp.Dial(amqpUrl)

	if err != nil {
		return nil, utils.ErrorHandler(err, "Error connecting to the CloudAMQP")
	}

	channel, err := connection.Channel()

	if err != nil {
		return nil, utils.ErrorHandler(err, "Error opening the channel")
	}

	if err := SetUpExchange(channel); err != nil {
		return nil, utils.ErrorHandler(err, err.Error())
	}

	if err := SetUpDlq(channel); err != nil {
		return nil, utils.ErrorHandler(err, err.Error())
	}

	if err := channel.Confirm(false); err != nil {
		return nil, utils.ErrorHandler(err, err.Error())
	}

	confirmations := channel.NotifyPublish(make(chan amqp.Confirmation, 100))

	return &AMQPConfig{Channel: channel, ConfirmationChannel: confirmations}, nil

}
