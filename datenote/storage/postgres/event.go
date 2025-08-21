package postgres

import (
	"context"
	"datenote/datenote/storage"
)

const insertEvent = `
	INSERT INTO events(
		name,
		date,
		info,
		category
	) VALUES(
		:name,
		:date,
		:info,
		:category
	)
	RETURNING id;
`
const getEvent = `
	SELECT * FROM events
	WHERE id=$1;
`
const getAllEvents = `
	SELECT * FROM events;
`
const updateEvent = `
	UPDATE events 
	SET name=:name,
	date=:date,
	info=:info,
	category=:category
	WHERE id=:id
	RETURNING id;
`
const deleteEvent = `
	DELETE FROM events 
	WHERE id=$1;
`
const searchEvent = `
	SELECT * FROM events 
	WHERE name ILIKE '%%' || $1 || '%%';
`

func (s *Storage) CreateEvent(ctx context.Context, t storage.Event) (int64, error) {
	stmt, err := s.db.PrepareNamed(insertEvent)
	if err != nil {
		return 0, err
	}

	var id int64
	if err := stmt.Get(&id, t); err != nil {
		return 0, err
	}
	return id, nil
}

func (s *Storage) GetEvent(ctx context.Context, id int64) (storage.Event, error) {
	var event storage.Event
	if err := s.db.Get(&event, getEvent, id); err != nil {
		return storage.Event{}, err
	}
	return event, nil
}

func (s *Storage) GetAllEvents(ctx context.Context) ([]storage.Event, error) {
	var event []storage.Event
	if err := s.db.Select(&event, getAllEvents); err != nil {
		return []storage.Event{}, err
	}
	return event, nil
}

func (s *Storage) UpdateEvent(ctx context.Context, t storage.Event) error {
	stmt, err := s.db.PrepareNamed(updateEvent)
	if err != nil {
		return err
	}

	var id int64
	if err := stmt.Get(&id, t); err != nil {
		return err
	}
	return nil
}

func (s *Storage) DeleteEvent(ctx context.Context, id int64) error {
	res := s.db.MustExec(deleteEvent, id)
	ok, err := res.RowsAffected()
	if err != nil || ok == 0 {
		return err
	}
	return nil
}

func (s *Storage) SearchEvent(ctx context.Context, sq string) ([]storage.Event, error) {
	var event []storage.Event
	if err := s.db.Select(&event, searchEvent, sq); err != nil {
		return []storage.Event{}, err
	}
	return event, nil
}
