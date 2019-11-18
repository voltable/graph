package expressions

import "errors"

var (
	MustBeReducible = errors.New("must be reducible")
	MustReduceToDifferent = errors.New("must reduce to different")
	ReducedNotCompatible = errors.New("reduced not compatible")
	ReducibleMustOverrideReduce = errors.New("reducible must override reduce")
	ArgumentCannotBeOfTypeVoid = errors.New("argument cannot be nil")
	ArgumentTypesMustMatch = errors.New("argument types must match")
	ArgumentMustBeBoolean = errors.New("argument must be boolean")
	MustRewriteChildToSameType = errors.New("must rewrite child to same type")
	ArgumentTypesMustBeLambda = errors.New("argument types must be lambda")
	InvalidBinaryOperations = errors.New("invalid binary operations")
)
