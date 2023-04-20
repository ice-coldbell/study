# CockroachDB

### 공부할 때 봐야하는 것

- [모든 것은 문서부터](https://www.cockroachlabs.com/docs/stable/)
    - [서버에게 SQL 튜닝은 필수](https://www.cockroachlabs.com/docs/v22.2/performance-best-practices-overview)
        - For `INSERT`, `UPSERT`, and `DELETE` statements, a single multi-row statement is faster than multiple single-row statements. Whenever possible, use multi-row statements for DML queries instead of multiple single-row statements. 요런거~
    - [DB의 Primary Key는 중요하지](https://www.cockroachlabs.com/docs/v22.2/schema-design-table.html#select-primary-key-columns)
    - [Index도 중요하지..](https://www.cockroachlabs.com/docs/v22.2/schema-design-indexes.html)
        - [hash-sharded-indexes](https://www.cockroachlabs.com/docs/v22.2/hash-sharded-indexes)
            - [https://www.cockroachlabs.com/blog/hash-sharded-indexes-unlock-linear-scaling-for-sequential-workloads/](https://www.cockroachlabs.com/blog/hash-sharded-indexes-unlock-linear-scaling-for-sequential-workloads/)
    - [Multi region](https://www.cockroachlabs.com/docs/v22.2/migrate-to-multiregion-sql)
    - [파티셔닝](https://www.cockroachlabs.com/docs/v22.2/partitioning.html)
- [딥하고, Low Level것은 코드로](https://github.com/cockroachdb/cockroach)

- 데브시스터즈에서 사용한다구!
    - [https://tech.devsisters.com/posts/cockroachdb-in-production/](https://tech.devsisters.com/posts/cockroachdb-in-production/) (cockroachdb 사용)
    - [https://tech.devsisters.com/posts/crk-launch-storage-postmortem/](https://tech.devsisters.com/posts/crk-launch-storage-postmortem/) (장애 과정 및 대응)
    - [https://tech.devsisters.com/posts/bit-level-database-hacking/](https://tech.devsisters.com/posts/bit-level-database-hacking/#ft_2) (장애복구방법)
    - [https://tech.devsisters.com/posts/crk-aws-az-failure-postmortem/](https://tech.devsisters.com/posts/crk-aws-az-failure-postmortem/) (az 장)

# ****CockroachDB란?****
간단히 살펴보자 자세한 건 문서로

- 분산형 SQL 데이터베이스로, 높은 확장성, 고가용성, 지리적 분산을 특징으로 합니다. 공개됨과 동시에 큰 관심을 받은 오픈 소스 프로젝트이며, NewSQL 데이터베이스 중 하나로 분류됩니다. Google Spanner의 영향을 받은 CockroachDB는 기업들이 대규모의 글로벌 애플리케이션을 구축하는 데 도움이 되도록 설계되었습니다.

```ruby
1. 글로벌에서 많은 트레픽을 받아야 할 때 사용된다.
2. NewSQL의 특징들이 필요한 경우에 사용된다.
```

## 주요 특징

- 확장성: CockroachDB는 수평적 확장성이 뛰어나 데이터와 트래픽이 증가함에 따라 쉽게 노드를 추가할 수 있습니다. 이를 통해 사용자는 대량의 데이터를 효율적으로 처리할 수 있습니다.
- 고가용성: 데이터를 여러 노드에 복제하여 저장함으로써 CockroachDB는 장애 발생 시 자동 복구 기능을 제공합니다. 이를 통해 데이터 손실을 최소화하고 서비스를 지속적으로 운영할 수 있습니다.
- 지리적 분산: CockroachDB는 데이터를 여러 지역이나 클라우드 환경에 분산 저장할 수 있습니다. 이를 통해 데이터의 지연 시간을 줄이고, 법적 규정을 준수하며, 지역 재해에 대비할 수 있습니다.
- 강력한 일관성: CockroachDB는 분산 환경에서도 강력한 일관성을 제공합니다. 사용자는 분산 트랜잭션을 실행하면서 일관된 결과를 얻을 수 있으며, ACID 속성이 유지됩니다.
- SQL 지원: CockroachDB는 PostgreSQL와 호환되는 SQL 쿼리를 지원합니다. 이로 인해 개발자들이 친숙한 SQL 쿼리 언어를 사용하여 데이터를 처리할 수 있습니다.