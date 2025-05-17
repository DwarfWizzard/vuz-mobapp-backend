package repository

import (
	"time"

	"github.com/DwarfWizzard/vuz-mobapp-backend/internal/education/dto"
)

func (r *Repo) ListSchedule(groupId uint32, date time.Time, isEvenWeek bool) ([]dto.Schedule, error) {
	var schedules []dto.Schedule

	var parity string
	if isEvenWeek {
		parity = "even"
	} else {
		parity = "odd"
	}

	err := r.dbClient.DB().Debug().Raw(
		`SELECT
			t.id template_id,
			COALESCE(o.date, CURRENT_TIMESTAMP) date,
			t.time_start,
			t.time_end,
			t.week_day,
			COALESCE(o.canceled, false) canceled,
			t.group_id,
			g.number group_number,
			subj.id subject_id,
			subj.name subject_name,
			teach.id teacher_id,
			teach.name teacher_name,
			loc.id location_id,
			TRIM(
    			loc.address ||
    			CASE WHEN loc.building IS NOT NULL AND loc.building <> '' THEN ', корп. ' || loc.building ELSE '' END ||
    			CASE WHEN loc.level IS NOT NULL AND loc.level <> '' THEN ', этаж ' || loc.level ELSE '' END ||
    			CASE WHEN loc.room IS NOT NULL AND loc.room <> '' THEN ', ауд. ' || loc.room ELSE '' END
			) AS address,
			o.note
		FROM schedule_templates t
		LEFT JOIN schedule_overrides o ON o.template_id = t.id AND o.date = $1
		JOIN edu_groups g ON g.id = t.group_id
		JOIN subjects subj ON subj.id = COALESCE(o.subject_id, t.subject_id)
		JOIN teachers teach ON teach.id = COALESCE(o.teacher_id, t.teacher_id)
		JOIN locations loc ON loc.id = COALESCE(o.location_id, t.location_id)
		WHERE
    		t.group_id = $2 AND
    		(t.week_parity = 'both' OR t.week_parity = $3)`,
		date, groupId, parity,
	).Scan(&schedules).Error

	if err != nil {
		return nil, err
	}

	return schedules, nil
}
