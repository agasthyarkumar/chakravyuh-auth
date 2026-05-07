echo "======================================="
echo "1. SUPERADMIN /me"
echo "======================================="

curl http://localhost:8080/me \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NzgxNDkzOTYsInJvbGUiOiJzdXBlcmFkbWluIiwidGVuYW50X2lkIjowLCJ1c2VyX2lkIjo1LCJ1c2VybmFtZSI6InN1cGVyYWRtaW4ifQ.6yFmYilDNW7dA2JFGZS5_4m6tv6xccdlBO5woqQmsL8"

echo
echo
echo "======================================="
echo "2. ADMIN /me"
echo "======================================="

curl http://localhost:8080/me \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NzgxNDk1NDEsInJvbGUiOiJhZG1pbiIsInRlbmFudF9pZCI6MSwidXNlcl9pZCI6MSwidXNlcm5hbWUiOiJhZ2FzdGh5YSJ9.xzCWys9zxfSAU6-R9G9jRrXrm0j-ms6ikKUhgEmE6QU"

echo
echo
echo "======================================="
echo "3. USER /me"
echo "======================================="

curl http://localhost:8080/me \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NzgxNDk1NjEsInJvbGUiOiJ1c2VyIiwidGVuYW50X2lkIjoxLCJ1c2VyX2lkIjozLCJ1c2VybmFtZSI6ImpvaG4ifQ.ws9-nYaSBloQdxs-pHLLE-QyExRrjiPBrLWfoYscFcw"

echo
echo
echo "======================================="
echo "4. ADMIN LIST USERS"
echo "======================================="

curl http://localhost:8080/admin/users \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NzgxNDk1NDEsInJvbGUiOiJhZG1pbiIsInRlbmFudF9pZCI6MSwidXNlcl9pZCI6MSwidXNlcm5hbWUiOiJhZ2FzdGh5YSJ9.xzCWys9zxfSAU6-R9G9jRrXrm0j-ms6ikKUhgEmE6QU"

echo
echo
echo "======================================="
echo "5. USER TRY ADMIN ROUTE (SHOULD FAIL)"
echo "======================================="

curl http://localhost:8080/admin/users \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NzgxNDk1NjEsInJvbGUiOiJ1c2VyIiwidGVuYW50X2lkIjoxLCJ1c2VyX2lkIjozLCJ1c2VybmFtZSI6ImpvaG4ifQ.ws9-nYaSBloQdxs-pHLLE-QyExRrjiPBrLWfoYscFcw"

echo
echo
echo "======================================="
echo "6. SUPERADMIN VIEW PENDING TENANTS"
echo "======================================="

curl http://localhost:8080/superadmin/tenants/pending \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NzgxNDkzOTYsInJvbGUiOiJzdXBlcmFkbWluIiwidGVuYW50X2lkIjowLCJ1c2VyX2lkIjo1LCJ1c2VybmFtZSI6InN1cGVyYWRtaW4ifQ.6yFmYilDNW7dA2JFGZS5_4m6tv6xccdlBO5woqQmsL8"

echo
echo
echo "======================================="
echo "7. ADMIN TRY SUPERADMIN ROUTE (SHOULD FAIL)"
echo "======================================="

curl http://localhost:8080/superadmin/tenants/pending \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NzgxNDk1NDEsInJvbGUiOiJhZG1pbiIsInRlbmFudF9pZCI6MSwidXNlcl9pZCI6MSwidXNlcm5hbWUiOiJhZ2FzdGh5YSJ9.xzCWys9zxfSAU6-R9G9jRrXrm0j-ms6ikKUhgEmE6QU"

echo
echo
echo "======================================="
echo "8. CREATE INVITATION"
echo "======================================="

curl -X POST http://localhost:8080/admin/invitations \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NzgxNDk1NDEsInJvbGUiOiJhZG1pbiIsInRlbmFudF9pZCI6MSwidXNlcl9pZCI6MSwidXNlcm5hbWUiOiJhZ2FzdGh5YSJ9.xzCWys9zxfSAU6-R9G9jRrXrm0j-ms6ikKUhgEmE6QU" \
-H "Content-Type: application/json" \
-d '{
  "email":"invite@test.com",
  "role":"user"
}'

echo
echo
echo "======================================="
echo "9. REFRESH TOKEN"
echo "======================================="

curl -X POST http://localhost:8080/refresh \
-H "Content-Type: application/json" \
-d '{
  "refresh_token":"33e683a96e1a9272e132a1f8762e6285f5d8a45122d4b91c1547d57acae1f359"
}'

echo
echo
echo "======================================="
echo "10. LOGOUT"
echo "======================================="

curl -X POST http://localhost:8080/logout \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NzgxNDk1NDEsInJvbGUiOiJhZG1pbiIsInRlbmFudF9pZCI6MSwidXNlcl9pZCI6MSwidXNlcm5hbWUiOiJhZ2FzdGh5YSJ9.xzCWys9zxfSAU6-R9G9jRrXrm0j-ms6ikKUhgEmE6QU" \
-H "Content-Type: application/json" \
-d '{
  "refresh_token":"33e683a96e1a9272e132a1f8762e6285f5d8a45122d4b91c1547d57acae1f359"
}'

echo
echo
echo "======================================="
echo "11. REFRESH AFTER LOGOUT (SHOULD FAIL)"
echo "======================================="

curl -X POST http://localhost:8080/refresh \
-H "Content-Type: application/json" \
-d '{
  "refresh_token":"33e683a96e1a9272e132a1f8762e6285f5d8a45122d4b91c1547d57acae1f359"
}'

echo ""