package constants

const (
	QueryInsertTraining           = "INSERT INTO Trainings(date, player_id, shooting_power, time, distance, successful_passes) VALUES(?,?,?,?,?,?)"
	QueryCountWeekPlayers         = "SELECT COUNT(*) AS total_players FROM ( SELECT player_id FROM Trainings WHERE WEEK(date, 3) = WEEK(CURDATE(), 3) GROUP BY player_id HAVING COUNT(*) >= ? ) AS subquery"
	QueryTotalValuesPerPlayerWeek = "SELECT player_id, p.name, SUM(shooting_power) AS total_shooting_power, SUM(time) AS total_time, SUM(distance) AS total_distance, SUM(successful_passes) AS total_successful_passes FROM Trainings JOIN Players p ON player_id = p.id WHERE WEEK(date, 3) = WEEK(CURDATE(), 3) GROUP BY player_id, p.name HAVING COUNT(*) >= 3"
	QueryCountPlayerByID          = "SELECT COUNT(*) FROM Players WHERE id = ?"
	QueryInsertPlayer             = "INSERT INTO Players (id, name, power, speed_distance, speed_time, passes, position) VALUES (?, ?, ?, ?, ?, ?, ?)"
	QueryUpdatePlayer             = "UPDATE Players SET name = ?, power = ?, speed_distance = ?, speed_time = ?, passes = ? WHERE id = ?"
	QueryDeletePlayerByID         = "DELETE FROM Players WHERE id = ?"
	QueryGetAllPlayers            = "SELECT id, name, power, speed_distance, speed_time, passes, position FROM Players"
	QueryGetConfiguration         = "SELECT * FROM Configuration LIMIT 1"
)
