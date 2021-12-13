package cleanCodeTips

// tips:
// 1. The more specific the function, the more general its name.

// func Parse(filepath string) (Config, error) {
//     switch fileExtension(filepath) {
//     case "json":
//         return parseJSON(filepath)
//     case "yaml":
//         return parseYAML(filepath)
//     case "toml":
//         return parseTOML(filepath)
//     default:
//         return Config{}, ErrUnknownFileExtension
//     }
// }

// 2. If the value, err := pattern is repeated more than once in a function,
//	  this is an indication that we can split the logic of our code into smaller pieces

// 3. Function signatures should only contain one or two input parameters.

// 4. our functions should only be 5–8 lines long, this can seem quite extreme at first.
//	  However, I feel that this rule is much easier to justify.
//    it is recommended to replace input parameters with an 'Options' struct instead.

// type QueueOptions struct {
//     Name string
//     Durable bool
//     DeleteOnExit bool
//     Exclusive bool
//     NoWait bool
//     Arguments []interface{}
// }

// q, err := ch.QueueDeclare(QueueOptions{
//     Name: "hello",
//     Durable: false,
//     DeleteOnExit: false,
//     Exclusive: false,
//     NoWait: false,
//     Arguments: nil,
// })

// 5. One final note on this: It's not always possible to change a function's signature.
//	  In this case, for example, we don't actually have control over our QueueDeclare
//	  function signature because it's from the RabbitMQ library.
//	  It's not our code, so we can't change it. However, we can wrap these functions to suit our purposes

// type RMQChannel struct {
//     channel *amqp.Channel
// }

// func (rmqch *RMQChannel) QueueDeclare(opts QueueOptions) (Queue, error) {
//     return rmqch.channel.QueueDeclare(
//         opts.Name,
//         opts.Durable,
//         opts.DeleteOnExit,
//         opts.Exclusive,
//         opts.NoWait,
//         opts.Arguments,
//     )
// }

// 6. Writing smaller functions can typically eliminate reliance on mutable variables that leak into the global scope.
