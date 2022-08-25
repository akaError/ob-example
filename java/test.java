import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.ResultSet;
import java.sql.SQLException;
import java.sql.Statement;

public class Test {
   public static void main(String[] args) {
       try {

            Class.forName("com.mysql.jdbc.Driver").newInstance();

            try{
                
                Connection connection = DriverManager.getConnection("jdbc:mysql://127.0.0.1:2881/test?user=root&password=");
                System.out.println(connection.getAutoCommit());
                Statement sm = connection.createStatement();
                String q1="drop table if exists t_meta_form";
                sm.executeUpdate(q1);
                String q2="CREATE TABLE t_meta_form ( name varchar(36) NOT NULL DEFAULT ' ', id int NOT NULL ) DEFAULT CHARSET = utf8mb4";
                String q3="insert into t_meta_form (name,id) values ('an','1')";
                sm.executeUpdate(q2);
                sm.executeUpdate(q3);                  

            }catch(SQLException se){
                System.out.println("error!");
                se.printStackTrace() ;
            }
            }catch (Exception ex) {
                ex.printStackTrace();
            }
    }
}


