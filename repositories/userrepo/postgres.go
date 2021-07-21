// Code generated by nero, DO NOT EDIT.
package userrepo

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/Masterminds/squirrel"
	"github.com/aneesh-jose/sample-todo/models"
	"github.com/pkg/errors"
	"github.com/sf9v/nero"
	"github.com/sf9v/nero/aggregate"
	"github.com/sf9v/nero/comparison"
	"github.com/sf9v/nero/sort"
)

// PostgresRepository is a repository that uses PostgreSQL as data store
type PostgresRepository struct {
	db     *sql.DB
	logger nero.Logger
	debug  bool
}

var _ Repository = (*PostgresRepository)(nil)

// NewPostgresRepository returns a PostresRepository
func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}

// Debug enables debug mode
func (pg *PostgresRepository) Debug() *PostgresRepository {
	return &PostgresRepository{
		db:     pg.db,
		debug:  true,
		logger: log.New(os.Stdout, "nero: ", 0),
	}
}

// WithLogger overrides the default logger
func (pg *PostgresRepository) WithLogger(logger nero.Logger) *PostgresRepository {
	pg.logger = logger
	return pg
}

// Tx begins a new transaction
func (pg *PostgresRepository) Tx(ctx context.Context) (nero.Tx, error) {
	return pg.db.BeginTx(ctx, nil)
}

// Create creates a User
func (pg *PostgresRepository) Create(ctx context.Context, c *Creator) (string, error) {
	return pg.create(ctx, pg.db, c)
}

// CreateTx creates a User in a transaction
func (pg *PostgresRepository) CreateTx(ctx context.Context, tx nero.Tx, c *Creator) (string, error) {
	txx, ok := tx.(*sql.Tx)
	if !ok {
		return "", errors.New("expecting tx to be *sql.Tx")
	}

	return pg.create(ctx, txx, c)
}

func (pg *PostgresRepository) create(ctx context.Context, runner nero.SQLRunner, c *Creator) (string, error) {
	if err := c.Validate(); err != nil {
		return "", err
	}

	columns := []string{
		"\"username\"",
		"\"name\"",
		"\"password\"",
	}

	values := []interface{}{
		c.username,
		c.name,
		c.password,
	}

	qb := squirrel.Insert("\"users\"").
		Columns(columns...).
		Values(values...).
		Suffix("RETURNING \"username\"").
		PlaceholderFormat(squirrel.Dollar).
		RunWith(runner)
	if pg.debug && pg.logger != nil {
		sql, args, err := qb.ToSql()
		pg.logger.Printf("method: Create, stmt: %q, args: %v, error: %v", sql, args, err)
	}

	var username string
	err := qb.QueryRowContext(ctx).Scan(&username)
	if err != nil {
		return "", err
	}

	return username, nil
}

// CreateMany batch creates Users
func (pg *PostgresRepository) CreateMany(ctx context.Context, cs ...*Creator) error {
	return pg.createMany(ctx, pg.db, cs...)
}

// CreateManyTx batch creates Users in a transaction
func (pg *PostgresRepository) CreateManyTx(ctx context.Context, tx nero.Tx, cs ...*Creator) error {
	txx, ok := tx.(*sql.Tx)
	if !ok {
		return errors.New("expecting tx to be *sql.Tx")
	}

	return pg.createMany(ctx, txx, cs...)
}

func (pg *PostgresRepository) createMany(ctx context.Context, runner nero.SQLRunner, cs ...*Creator) error {
	if len(cs) == 0 {
		return nil
	}

	columns := []string{
		"\"username\"",
		"\"name\"",
		"\"password\"",
	}
	qb := squirrel.Insert("\"users\"").Columns(columns...)
	for _, c := range cs {
		if err := c.Validate(); err != nil {
			return err
		}

		qb = qb.Values(
			c.username,
			c.name,
			c.password,
		)
	}

	qb = qb.Suffix("RETURNING \"username\"").
		PlaceholderFormat(squirrel.Dollar)
	if pg.debug && pg.logger != nil {
		sql, args, err := qb.ToSql()
		pg.logger.Printf("method: CreateMany, stmt: %q, args: %v, error: %v", sql, args, err)
	}

	_, err := qb.RunWith(runner).ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}

// Query queries Users
func (pg *PostgresRepository) Query(ctx context.Context, q *Queryer) ([]*models.User, error) {
	return pg.query(ctx, pg.db, q)
}

// QueryTx queries Users in a transaction
func (pg *PostgresRepository) QueryTx(ctx context.Context, tx nero.Tx, q *Queryer) ([]*models.User, error) {
	txx, ok := tx.(*sql.Tx)
	if !ok {
		return nil, errors.New("expecting tx to be *sql.Tx")
	}

	return pg.query(ctx, txx, q)
}

func (pg *PostgresRepository) query(ctx context.Context, runner nero.SQLRunner, q *Queryer) ([]*models.User, error) {
	qb := pg.buildSelect(q)
	if pg.debug && pg.logger != nil {
		sql, args, err := qb.ToSql()
		pg.logger.Printf("method: Query, stmt: %q, args: %v, error: %v", sql, args, err)
	}

	rows, err := qb.RunWith(runner).QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []*models.User{}
	for rows.Next() {
		var user models.User
		err = rows.Scan(
			&user.Username,
			&user.Name,
			&user.Password,
		)
		if err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	return users, nil
}

// QueryOne queries a User
func (pg *PostgresRepository) QueryOne(ctx context.Context, q *Queryer) (*models.User, error) {
	return pg.queryOne(ctx, pg.db, q)
}

// QueryOneTx queries a User in a transaction
func (pg *PostgresRepository) QueryOneTx(ctx context.Context, tx nero.Tx, q *Queryer) (*models.User, error) {
	txx, ok := tx.(*sql.Tx)
	if !ok {
		return nil, errors.New("expecting tx to be *sql.Tx")
	}

	return pg.queryOne(ctx, txx, q)
}

func (pg *PostgresRepository) queryOne(ctx context.Context, runner nero.SQLRunner, q *Queryer) (*models.User, error) {
	qb := pg.buildSelect(q)
	if pg.debug && pg.logger != nil {
		sql, args, err := qb.ToSql()
		pg.logger.Printf("method: QueryOne, stmt: %q, args: %v, error: %v", sql, args, err)
	}

	var user models.User
	err := qb.RunWith(runner).
		QueryRowContext(ctx).
		Scan(
			&user.Username,
			&user.Name,
			&user.Password,
		)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (pg *PostgresRepository) buildSelect(q *Queryer) squirrel.SelectBuilder {
	columns := []string{
		"\"username\"",
		"\"name\"",
		"\"password\"",
	}
	qb := squirrel.Select(columns...).
		From("\"users\"").
		PlaceholderFormat(squirrel.Dollar)

	pfs := q.pfs
	pb := &comparison.Predicates{}
	for _, pf := range pfs {
		pf(pb)
	}

	for _, p := range pb.All() {
		switch p.Op {
		case comparison.Eq:
			col, ok := p.Arg.(Column)
			if ok {
				qb = qb.Where(fmt.Sprintf("%q = %q", p.Col, col.String()))
			} else {
				qb = qb.Where(fmt.Sprintf("%q = ?", p.Col), p.Arg)
			}
		case comparison.NotEq:
			col, ok := p.Arg.(Column)
			if ok {
				qb = qb.Where(fmt.Sprintf("%q <> %q", p.Col, col.String()))
			} else {
				qb = qb.Where(fmt.Sprintf("%q <> ?", p.Col), p.Arg)
			}
		case comparison.Gt:
			col, ok := p.Arg.(Column)
			if ok {
				qb = qb.Where(fmt.Sprintf("%q > %q", p.Col, col.String()))
			} else {
				qb = qb.Where(fmt.Sprintf("%q > ?", p.Col), p.Arg)
			}
		case comparison.GtOrEq:
			col, ok := p.Arg.(Column)
			if ok {
				qb = qb.Where(fmt.Sprintf("%q >= %q", p.Col, col.String()))
			} else {
				qb = qb.Where(fmt.Sprintf("%q >= ?", p.Col), p.Arg)
			}
		case comparison.Lt:
			col, ok := p.Arg.(Column)
			if ok {
				qb = qb.Where(fmt.Sprintf("%q < %q", p.Col, col.String()))
			} else {
				qb = qb.Where(fmt.Sprintf("%q < ?", p.Col), p.Arg)
			}
		case comparison.LtOrEq:
			col, ok := p.Arg.(Column)
			if ok {
				qb = qb.Where(fmt.Sprintf("%q <= %q", p.Col, col.String()))
			} else {
				qb = qb.Where(fmt.Sprintf("%q <= ?", p.Col), p.Arg)
			}
		case comparison.IsNull:
			qb = qb.Where(fmt.Sprintf("%q IS NULL", p.Col))
		case comparison.IsNotNull:
			qb = qb.Where(fmt.Sprintf("%q IS NOT NULL", p.Col))
		case comparison.In, comparison.NotIn:
			args := p.Arg.([]interface{})
			if len(args) == 0 {
				continue
			}
			qms := []string{}
			for range args {
				qms = append(qms, "?")
			}
			fmtStr := "%q IN (%s)"
			if p.Op == comparison.NotIn {
				fmtStr = "%q NOT IN (%s)"
			}
			plchldr := strings.Join(qms, ",")
			qb = qb.Where(fmt.Sprintf(fmtStr, p.Col, plchldr), args...)
		}
	}

	sfs := q.sfs
	sorts := &sort.Sorts{}
	for _, sf := range sfs {
		sf(sorts)
	}
	for _, s := range sorts.All() {
		col := fmt.Sprintf("%q", s.Col)
		switch s.Direction {
		case sort.Asc:
			qb = qb.OrderBy(col + " ASC")
		case sort.Desc:
			qb = qb.OrderBy(col + " DESC")
		}
	}

	if q.limit > 0 {
		qb = qb.Limit(uint64(q.limit))
	}

	if q.offset > 0 {
		qb = qb.Offset(uint64(q.offset))
	}

	return qb
}

// Update updates a User or many Users
func (pg *PostgresRepository) Update(ctx context.Context, u *Updater) (int64, error) {
	return pg.update(ctx, pg.db, u)
}

// UpdateTx updates a User many Users in a transaction
func (pg *PostgresRepository) UpdateTx(ctx context.Context, tx nero.Tx, u *Updater) (int64, error) {
	txx, ok := tx.(*sql.Tx)
	if !ok {
		return 0, errors.New("expecting tx to be *sql.Tx")
	}

	return pg.update(ctx, txx, u)
}

func (pg *PostgresRepository) update(ctx context.Context, runner nero.SQLRunner, u *Updater) (int64, error) {
	qb := squirrel.Update("\"users\"").
		PlaceholderFormat(squirrel.Dollar)

	cnt := 0

	if !isZero(u.name) {
		qb = qb.Set("\"name\"", u.name)
		cnt++
	}

	if !isZero(u.password) {
		qb = qb.Set("\"password\"", u.password)
		cnt++
	}

	if cnt == 0 {
		return 0, nil
	}

	pfs := u.pfs
	pb := &comparison.Predicates{}
	for _, pf := range pfs {
		pf(pb)
	}

	for _, p := range pb.All() {
		switch p.Op {
		case comparison.Eq:
			col, ok := p.Arg.(Column)
			if ok {
				qb = qb.Where(fmt.Sprintf("%q = %q", p.Col, col.String()))
			} else {
				qb = qb.Where(fmt.Sprintf("%q = ?", p.Col), p.Arg)
			}
		case comparison.NotEq:
			col, ok := p.Arg.(Column)
			if ok {
				qb = qb.Where(fmt.Sprintf("%q <> %q", p.Col, col.String()))
			} else {
				qb = qb.Where(fmt.Sprintf("%q <> ?", p.Col), p.Arg)
			}
		case comparison.Gt:
			col, ok := p.Arg.(Column)
			if ok {
				qb = qb.Where(fmt.Sprintf("%q > %q", p.Col, col.String()))
			} else {
				qb = qb.Where(fmt.Sprintf("%q > ?", p.Col), p.Arg)
			}
		case comparison.GtOrEq:
			col, ok := p.Arg.(Column)
			if ok {
				qb = qb.Where(fmt.Sprintf("%q >= %q", p.Col, col.String()))
			} else {
				qb = qb.Where(fmt.Sprintf("%q >= ?", p.Col), p.Arg)
			}
		case comparison.Lt:
			col, ok := p.Arg.(Column)
			if ok {
				qb = qb.Where(fmt.Sprintf("%q < %q", p.Col, col.String()))
			} else {
				qb = qb.Where(fmt.Sprintf("%q < ?", p.Col), p.Arg)
			}
		case comparison.LtOrEq:
			col, ok := p.Arg.(Column)
			if ok {
				qb = qb.Where(fmt.Sprintf("%q <= %q", p.Col, col.String()))
			} else {
				qb = qb.Where(fmt.Sprintf("%q <= ?", p.Col), p.Arg)
			}
		case comparison.IsNull:
			qb = qb.Where(fmt.Sprintf("%q IS NULL", p.Col))
		case comparison.IsNotNull:
			qb = qb.Where(fmt.Sprintf("%q IS NOT NULL", p.Col))
		case comparison.In, comparison.NotIn:
			args := p.Arg.([]interface{})
			if len(args) == 0 {
				continue
			}
			qms := []string{}
			for range args {
				qms = append(qms, "?")
			}
			fmtStr := "%q IN (%s)"
			if p.Op == comparison.NotIn {
				fmtStr = "%q NOT IN (%s)"
			}
			plchldr := strings.Join(qms, ",")
			qb = qb.Where(fmt.Sprintf(fmtStr, p.Col, plchldr), args...)
		}
	}

	if pg.debug && pg.logger != nil {
		sql, args, err := qb.ToSql()
		pg.logger.Printf("method: Update, stmt: %q, args: %v, error: %v", sql, args, err)
	}

	res, err := qb.RunWith(runner).ExecContext(ctx)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

// Delete deletes a User or many Users
func (pg *PostgresRepository) Delete(ctx context.Context, d *Deleter) (int64, error) {
	return pg.delete(ctx, pg.db, d)
}

// Delete deletes a User or many Users in a transaction
func (pg *PostgresRepository) DeleteTx(ctx context.Context, tx nero.Tx, d *Deleter) (int64, error) {
	txx, ok := tx.(*sql.Tx)
	if !ok {
		return 0, errors.New("expecting tx to be *sql.Tx")
	}

	return pg.delete(ctx, txx, d)
}

func (pg *PostgresRepository) delete(ctx context.Context, runner nero.SQLRunner, d *Deleter) (int64, error) {
	qb := squirrel.Delete("\"users\"").
		PlaceholderFormat(squirrel.Dollar)

	pfs := d.pfs
	pb := &comparison.Predicates{}
	for _, pf := range pfs {
		pf(pb)
	}

	for _, p := range pb.All() {
		switch p.Op {
		case comparison.Eq:
			col, ok := p.Arg.(Column)
			if ok {
				qb = qb.Where(fmt.Sprintf("%q = %q", p.Col, col.String()))
			} else {
				qb = qb.Where(fmt.Sprintf("%q = ?", p.Col), p.Arg)
			}
		case comparison.NotEq:
			col, ok := p.Arg.(Column)
			if ok {
				qb = qb.Where(fmt.Sprintf("%q <> %q", p.Col, col.String()))
			} else {
				qb = qb.Where(fmt.Sprintf("%q <> ?", p.Col), p.Arg)
			}
		case comparison.Gt:
			col, ok := p.Arg.(Column)
			if ok {
				qb = qb.Where(fmt.Sprintf("%q > %q", p.Col, col.String()))
			} else {
				qb = qb.Where(fmt.Sprintf("%q > ?", p.Col), p.Arg)
			}
		case comparison.GtOrEq:
			col, ok := p.Arg.(Column)
			if ok {
				qb = qb.Where(fmt.Sprintf("%q >= %q", p.Col, col.String()))
			} else {
				qb = qb.Where(fmt.Sprintf("%q >= ?", p.Col), p.Arg)
			}
		case comparison.Lt:
			col, ok := p.Arg.(Column)
			if ok {
				qb = qb.Where(fmt.Sprintf("%q < %q", p.Col, col.String()))
			} else {
				qb = qb.Where(fmt.Sprintf("%q < ?", p.Col), p.Arg)
			}
		case comparison.LtOrEq:
			col, ok := p.Arg.(Column)
			if ok {
				qb = qb.Where(fmt.Sprintf("%q <= %q", p.Col, col.String()))
			} else {
				qb = qb.Where(fmt.Sprintf("%q <= ?", p.Col), p.Arg)
			}
		case comparison.IsNull:
			qb = qb.Where(fmt.Sprintf("%q IS NULL", p.Col))
		case comparison.IsNotNull:
			qb = qb.Where(fmt.Sprintf("%q IS NOT NULL", p.Col))
		case comparison.In, comparison.NotIn:
			args := p.Arg.([]interface{})
			if len(args) == 0 {
				continue
			}
			qms := []string{}
			for range args {
				qms = append(qms, "?")
			}
			fmtStr := "%q IN (%s)"
			if p.Op == comparison.NotIn {
				fmtStr = "%q NOT IN (%s)"
			}
			plchldr := strings.Join(qms, ",")
			qb = qb.Where(fmt.Sprintf(fmtStr, p.Col, plchldr), args...)
		}
	}

	if pg.debug && pg.logger != nil {
		sql, args, err := qb.ToSql()
		pg.logger.Printf("method: Delete, stmt: %q, args: %v, error: %v", sql, args, err)
	}

	res, err := qb.RunWith(runner).ExecContext(ctx)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

// Aggregate runs an aggregate query
func (pg *PostgresRepository) Aggregate(ctx context.Context, a *Aggregator) error {
	return pg.aggregate(ctx, pg.db, a)
}

// Aggregate runs an aggregate query in a transaction
func (pg *PostgresRepository) AggregateTx(ctx context.Context, tx nero.Tx, a *Aggregator) error {
	txx, ok := tx.(*sql.Tx)
	if !ok {
		return errors.New("expecting tx to be *sql.Tx")
	}

	return pg.aggregate(ctx, txx, a)
}

func (pg *PostgresRepository) aggregate(ctx context.Context, runner nero.SQLRunner, a *Aggregator) error {
	aggs := &aggregate.Aggregates{}
	for _, aggf := range a.aggfs {
		aggf(aggs)
	}
	cols := []string{}
	for _, agg := range aggs.All() {
		col := agg.Col
		qcol := fmt.Sprintf("%q", col)
		switch agg.Fn {
		case aggregate.Avg:
			cols = append(cols, "AVG("+qcol+") avg_"+col)
		case aggregate.Count:
			cols = append(cols, "COUNT("+qcol+") count_"+col)
		case aggregate.Max:
			cols = append(cols, "MAX("+qcol+") max_"+col)
		case aggregate.Min:
			cols = append(cols, "MIN("+qcol+") min_"+col)
		case aggregate.Sum:
			cols = append(cols, "SUM("+qcol+") sum_"+col)
		case aggregate.None:
			cols = append(cols, qcol)
		}
	}

	qb := squirrel.Select(cols...).From("\"users\"").
		PlaceholderFormat(squirrel.Dollar)

	groups := []string{}
	for _, group := range a.groups {
		groups = append(groups, fmt.Sprintf("%q", group.String()))
	}
	qb = qb.GroupBy(groups...)

	pfs := a.pfs
	pb := &comparison.Predicates{}
	for _, pf := range pfs {
		pf(pb)
	}

	for _, p := range pb.All() {
		switch p.Op {
		case comparison.Eq:
			col, ok := p.Arg.(Column)
			if ok {
				qb = qb.Where(fmt.Sprintf("%q = %q", p.Col, col.String()))
			} else {
				qb = qb.Where(fmt.Sprintf("%q = ?", p.Col), p.Arg)
			}
		case comparison.NotEq:
			col, ok := p.Arg.(Column)
			if ok {
				qb = qb.Where(fmt.Sprintf("%q <> %q", p.Col, col.String()))
			} else {
				qb = qb.Where(fmt.Sprintf("%q <> ?", p.Col), p.Arg)
			}
		case comparison.Gt:
			col, ok := p.Arg.(Column)
			if ok {
				qb = qb.Where(fmt.Sprintf("%q > %q", p.Col, col.String()))
			} else {
				qb = qb.Where(fmt.Sprintf("%q > ?", p.Col), p.Arg)
			}
		case comparison.GtOrEq:
			col, ok := p.Arg.(Column)
			if ok {
				qb = qb.Where(fmt.Sprintf("%q >= %q", p.Col, col.String()))
			} else {
				qb = qb.Where(fmt.Sprintf("%q >= ?", p.Col), p.Arg)
			}
		case comparison.Lt:
			col, ok := p.Arg.(Column)
			if ok {
				qb = qb.Where(fmt.Sprintf("%q < %q", p.Col, col.String()))
			} else {
				qb = qb.Where(fmt.Sprintf("%q < ?", p.Col), p.Arg)
			}
		case comparison.LtOrEq:
			col, ok := p.Arg.(Column)
			if ok {
				qb = qb.Where(fmt.Sprintf("%q <= %q", p.Col, col.String()))
			} else {
				qb = qb.Where(fmt.Sprintf("%q <= ?", p.Col), p.Arg)
			}
		case comparison.IsNull:
			qb = qb.Where(fmt.Sprintf("%q IS NULL", p.Col))
		case comparison.IsNotNull:
			qb = qb.Where(fmt.Sprintf("%q IS NOT NULL", p.Col))
		case comparison.In, comparison.NotIn:
			args := p.Arg.([]interface{})
			if len(args) == 0 {
				continue
			}
			qms := []string{}
			for range args {
				qms = append(qms, "?")
			}
			fmtStr := "%q IN (%s)"
			if p.Op == comparison.NotIn {
				fmtStr = "%q NOT IN (%s)"
			}
			plchldr := strings.Join(qms, ",")
			qb = qb.Where(fmt.Sprintf(fmtStr, p.Col, plchldr), args...)
		}
	}

	sfs := a.sfs
	sorts := &sort.Sorts{}
	for _, sf := range sfs {
		sf(sorts)
	}
	for _, s := range sorts.All() {
		col := fmt.Sprintf("%q", s.Col)
		switch s.Direction {
		case sort.Asc:
			qb = qb.OrderBy(col + " ASC")
		case sort.Desc:
			qb = qb.OrderBy(col + " DESC")
		}
	}

	if pg.debug && pg.logger != nil {
		sql, args, err := qb.ToSql()
		pg.logger.Printf("method: Aggregate, stmt: %q, args: %v, error: %v", sql, args, err)
	}

	rows, err := qb.RunWith(runner).QueryContext(ctx)
	if err != nil {
		return err
	}
	defer rows.Close()

	v := reflect.ValueOf(a.v).Elem()
	t := reflect.TypeOf(v.Interface()).Elem()
	if t.NumField() != len(cols) {
		return errors.New("aggregate columns and destination struct field count should match")
	}

	for rows.Next() {
		ve := reflect.New(t).Elem()
		dest := make([]interface{}, ve.NumField())
		for i := 0; i < ve.NumField(); i++ {
			dest[i] = ve.Field(i).Addr().Interface()
		}

		err = rows.Scan(dest...)
		if err != nil {
			return err
		}

		v.Set(reflect.Append(v, ve))
	}

	return nil
}
